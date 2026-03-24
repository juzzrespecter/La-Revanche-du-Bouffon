package crawler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"slices"
	"spider/internal/client"
	hp "spider/internal/htmlparse"
	"spider/internal/logger"
	log "spider/internal/logger"
	"spider/internal/utils"
	"sync"
	//borrame
)

type Config struct {
	IsRecursive bool
	Depth       uint
	StoreDir    string
}

func (c Config) Unpack() (bool, uint, string) {
	return c.IsRecursive, c.Depth, c.StoreDir
}

func fetchImages(src []string, storage string, c *http.Client) {
	e := make(chan error)

	var wg sync.WaitGroup

	for _, url := range src {
		wg.Go(func() {
			req, err := utils.GenerateRequest(url)
			res, err := c.Do(req)
			if err != nil {
				e <- err
				return
			}
			defer res.Body.Close()
			filePath := storage + path.Base(url)
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

		})
	}
}

func fetchUrls(urls []string, c *client.CustomClient) ([]string, []string) {
	r := make(chan hp.ParseResult, len(urls))
	e := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
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
			res, err := c.Get(u)
			if err != nil {
				e <- err
				return
			}
			logger.Info(fmt.Sprintf("Request to %s %d", u, res.StatusCode))
			refs, err := hp.ParseHtml(res.Body)
			if err != nil {
				e <- fmt.Errorf("Invalid body")
				return
			}
			for i, ref := range refs.Href {
				refs.Href[i] = utils.SetUpURL(urlData, ref)
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

func Crawl(url url.URL, cfg Config) error {
	c := client.NewClient()

	isRecursive, depth, storage := cfg.Unpack()
	urls := []string{url.String()}
	var recursiveCrawl func([]string, uint)
	recursiveCrawl = func(urls []string, lvl uint) {
		hrefs, srcs := fetchUrls(urls, c)
		fetchImages(srcs, storage, c)
		logger.Debug(fmt.Sprintf("isRecursive: %d, depth: %d, lvl: %d, len: %d\n", isRecursive, depth, lvl, len(hrefs)))
		if isRecursive && depth > lvl && len(hrefs) > 0 {
			recursiveCrawl(hrefs, lvl+1)
		}
	}
	recursiveCrawl(urls, 1)

	// limitar concurrencias
	// gestion de timeout
	// gestion de retries
	// gestion de redirecciones ??
	// gestion de valores de respuesta
	// gestion de urls ya visitadas
	return nil
}
