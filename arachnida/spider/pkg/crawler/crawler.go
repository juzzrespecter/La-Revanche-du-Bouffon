package crawler

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"slices"
	"spider/internal/client"
	hp "spider/internal/htmlparse"
	"spider/internal/logger"
	log "spider/internal/logger"
	"spider/internal/utils"
	"sync"
	"time"
)

type Config struct {
	Ctx                   context.Context
	IsRecursive           bool
	Depth                 uint
	StoreDir              string
	Timeout               time.Duration
	MaxConcurrentRequests uint
}

func (c Config) Unpack() (bool, uint, string) {
	return c.IsRecursive, c.Depth, c.StoreDir
}

func fetchImages(src []string, storage string, c *client.CustomClient) {
	e := make(chan error)

	var wg sync.WaitGroup

	for _, url := range src {
		wg.Go(func() {
			if c.AlreadyVisited(url) {
				return
			}
			res, cancel, err := c.Get(url)
			defer cancel()
			if err != nil {
				e <- err
				return
			}
			defer res.Body.Close()
			filePath := utils.GenerateFileName(storage, url)
			f, err := os.Create(filePath)
			if err != nil {
				e <- err
				return
			}
			defer f.Close()
			if _, err = io.Copy(f, res.Body); err != nil {
				e <- err
				return
			}
			info := fmt.Sprintf("%s: saved successfully", filePath)
			log.Info(info)
		})
	}
	go func() {
		defer close(e)
		wg.Wait()
	}()
	for err := range e {
		if err != nil {
			log.Warning(err.Error())
		}
	}
}

func fetchUrls(urls []string, c *client.CustomClient) ([]string, []string) {
	r := make(chan hp.ParseResult, len(urls))
	e := make(chan error)
	ctx, cancel := context.WithCancel(c.Ctx)
	hrefs := make([]string, 0)
	srcs := make([]string, 0)

	wg := &sync.WaitGroup{}
	for _, u := range urls {
		wg.Go(func() {
			urlData, err := url.Parse(u)
			if err != nil {
				e <- err
				return
			}
			if c.AlreadyVisited(u) {
				return
			}
			res, cancel, err := c.Get(u)
			defer cancel()
			if err != nil {
				e <- err
				return
			}
			log.Info(fmt.Sprintf("Request to %s %d", u, res.StatusCode))
			refs, err := hp.ParseHtml(res.Body, c.HostBound)
			if err != nil {
				e <- fmt.Errorf("%s: %w", u, err)
				return
			}

			for i, ref := range refs.Href {
				refs.Href[i] = utils.SetUpURL(urlData, ref)
			}
			for i, src := range refs.Src {
				refs.Src[i] = utils.SetUpURL(urlData, src)
			}
			r <- refs
		})
	}
	go func() {
		wg.Wait()
		close(r)
		close(e)
		cancel()
	}()

WaitRoutines:
	for {
		select {
		case x := <-r:
			href, src := x.Unpack()
			hrefs = append(hrefs, href...)
			srcs = append(srcs, src...)
		case y := <-e:
			if y != nil {
				log.Warning(y.Error())
			}
		case <-ctx.Done():
			break WaitRoutines
		}
	}
	hrefs = slices.Compact(hrefs)
	srcs = slices.Compact(srcs)
	return hrefs, srcs
}

func Crawl(url url.URL, cfg *Config) error {
	c := client.NewClient(
		cfg.Ctx,
		url.Hostname(),
		cfg.Timeout,
		cfg.MaxConcurrentRequests,
	)
	defer c.CloseRequestSemaphore()
	isRecursive, depth, storage := cfg.Unpack()
	urls := []string{url.String()}
	var recursiveCrawl func([]string, uint)
	recursiveCrawl = func(urls []string, lvl uint) {
		hrefs, srcs := fetchUrls(urls, c)
		fetchImages(srcs, storage, c)
		if isRecursive && depth > lvl && len(hrefs) > 0 {
			logger.Info(fmt.Sprintf("Trying depth level: %d...\n", lvl))
			recursiveCrawl(hrefs, lvl+1)
		}
	}
	recursiveCrawl(urls, 1)
	log.Info("Done.")
	return nil
}
