package client

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"spider/internal/logger"
	"spider/internal/utils"
	"sync"
	"time"
)

type CustomClient struct {
	client      *http.Client
	visitedUrls map[string]bool
	m           sync.RWMutex
	maxRetries  int
}

const retriableCodes = []int{
	http.StatusRequestTimeout,
	http.StatusTooEarly,
	http.StatusTooManyRequests,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
}

func backoff(retryAttempt int) {

}

func NewClient(ctx context.Context) *CustomClient {
	return &CustomClient{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		visitedUrls: make(map[string]bool),
		maxRetries:  3,
	}
}

func (c *CustomClient) alreadyVisited(url string) bool {
	c.m.RLock()
	status := c.visitedUrls[url]
	if status == false {
		c.visitedUrls[url] = true
	}
	c.m.RUnlock()
	return status
}

/*
Criterios para retry:

  - http handler error: no

  - 408, 425, 429: si

  - 500, 502, 503, 504: si

    https://www.restapitutorial.com/advanced/responses/retries

Exponential Backoffs y retries:

	Los delays fijos nos valen pinga si las gorutinas de peticiones
	se despiertan a la vez para hacer el retry.

	sleep = jitter(min(cap, base × factor^attempt))
*/
func (c *CustomClient) Get(url string) (*http.Response, error) {
	if c.alreadyVisited(url) {
		logger.Debug(fmt.Sprintf("%s: already visited\n", url))
		return nil, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", utils.GetUserAgent())
	for retry := 1; retry < c.maxRetries; retry++ {
		req := req.Clone(req.Context())
		res, err := c.client.Do(req)
		switch {
		case err != nil:
			return nil, err
		case slices.Contains(retriableCodes, res.StatusCode):
			select {
			case <-time.After(time.Second):
			case <-c.ctx.Done():
				return nil, c.ctx.Err()
			}
		case res.StatusCode == http.StatusOK:
			return res, nil
		default:
			return nil, fmt.Errorf("Request to %s %d", url, res.StatusCode)
		}
	}
}

/*


req, err := utils.GenerateRequest(u)
			if err != nil {
				e <- err
				return
			}
			res, err := c.Do(req)
			// manejar 429 y hacerlo generico para pillar srcs
			switch {
			case err != nil:
				e <- err
				return
			case res.StatusCode > http.StatusBadRequest:
				e <- fmt.Errorf("Request to %s %d", u, res.StatusCode)
				return
			}
*/
