package client

import (
	"context"
	"fmt"
	"math/rand"
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

	Ctx *context.Context
}

var retriableCodes = []int{
	http.StatusRequestTimeout,
	http.StatusTooEarly,
	http.StatusTooManyRequests,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
}

func backoff(retryAttempt int) time.Duration {
	base := 500 * time.Millisecond
	max := 5 * time.Second

	x := time.Duration(1<<retryAttempt) * base
	if x > max {
		x = max
	}
	y := int64(x / 2)
	return time.Duration(rand.Int63n(int64(x)-y) + y)
}

func NewClient(timeout time.Duration, ctx context.Context) *CustomClient {
	return &CustomClient{
		client: &http.Client{
			Timeout: timeout,
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
	ctx, cancel := context.WithTimeout(*c.Ctx, 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
			retryTime := backoff(retry)
			select {
			case <-time.After(retryTime):
				logger.Warning(fmt.Sprintf("Retrying for %s (attempt %d)...\n", url, retry))
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		case res.StatusCode == http.StatusOK:
			return res, nil
		default:
			return nil, fmt.Errorf("Request to %s %d", url, res.StatusCode)
		}
	}
	return nil, fmt.Errorf("%s: max number of retries exceeded", url)
}
