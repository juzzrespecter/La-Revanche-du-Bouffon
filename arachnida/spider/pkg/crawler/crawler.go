package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"spider/internal/htmlparse"
	"sync"
)

func Crawl(url url.URL) error {
	var c http.Client
	ch := make(chan int, 500)
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Go(func() {
			res, err := c.Get(url.String())
			if err != nil {
				fmt.Println(err)
			}
			ch <- i
		})
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var output []int
	for i := range ch {
		fmt.Println("Waiting channel...")
		output = append(output, i)
	}
	// gestion de user agents
	// gestion de timeout
	// gestion de retries
	// gestion de redirecciones
	// gestion de valores de respuesta
	res, err := req.Get("http://localhost:8082")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	src, href, err := htmlparse.ParseHtml(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(src, href)
	return nil
}
