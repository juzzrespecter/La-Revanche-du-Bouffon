package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"spider/internal/htmlparse"
	"sync"
)

func Crawl(url url.URL) {
	var req http.Client
	ch := make(chan int, 500)
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Called goroutine %d\n", i)
			ch <- i
		}()
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
	html, err := io.ReadAll(res.Body)
	htmlparse.ParseHtml(res.Body)
}
