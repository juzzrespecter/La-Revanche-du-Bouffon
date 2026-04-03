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
	m           sync.Mutex
	maxRetries  int
	reqSem      chan struct{}

	Ctx                   context.Context
	HostBound             string
	Timeout               time.Duration
	MaxConcurrentRequests uint
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
	base := 3000 * time.Millisecond
	max := 10 * time.Second

	x := time.Duration(1<<retryAttempt) * base
	x = min(x, max)
	y := int64(x / 2)
	return time.Duration(rand.Int63n(int64(x)-y) + y)
}

func NewClient(
	ctx context.Context,
	host string,
	timeout time.Duration,
	maxConcurrentRequests uint) *CustomClient {
	n := &CustomClient{
		client: &http.Client{
			Timeout: timeout,
		},
		visitedUrls: make(map[string]bool),
		maxRetries:  3,

		Ctx:                   ctx,
		HostBound:             host,
		Timeout:               timeout,
		MaxConcurrentRequests: maxConcurrentRequests,
		reqSem:                make(chan struct{}, maxConcurrentRequests),
	}
	return n
}

func (c *CustomClient) CloseRequestSemaphore() {
	close(c.reqSem)
}

func (c *CustomClient) AlreadyVisited(url string) bool {
	c.m.Lock()
	status := c.visitedUrls[url]
	if status == false {
		c.visitedUrls[url] = true
	}
	c.m.Unlock()
	return status
}

func (c *CustomClient) freeSem() {
	select {
	case <-c.reqSem:
	default:
		return
	}
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
func (c *CustomClient) Get(url string) (*http.Response, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(c.Ctx, c.Timeout)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, cancel, err
	}
	defer func() {
		c.freeSem()
	}()
	select {
	case c.reqSem <- struct{}{}:
		for retry := 0; retry < c.maxRetries; retry++ {
			if retry > 0 {
				req = req.Clone(req.Context())
			}
			req.Header.Set("User-Agent", utils.GetUserAgent())
			logger.Info("Trying for " + url + " ...")
			res, err := c.client.Do(req)
			switch {
			case err != nil:
				return nil, cancel, fmt.Errorf("%s: %s", url, err)
			case slices.Contains(retriableCodes, res.StatusCode):
				if req.Body != nil {
					req.Body.Close()
				}
				select {
				case <-time.After(backoff(retry)):
					logger.Warning(fmt.Sprintf("Retrying for %s (status: %d) (attempt %d)...\n", url, res.StatusCode, retry+1))
				case <-ctx.Done():
					return nil, cancel, fmt.Errorf("%s: %s", url, err)
				}
			case res.StatusCode == http.StatusOK:
				if retry > 0 {
					fmt.Println("ueeeeeevamoo")
				}
				return res, cancel, nil
			default:
				return nil, cancel, fmt.Errorf("Request to %s failed with %d", url, res.StatusCode)
			}
		}
	case <-ctx.Done():
		return nil, cancel, ctx.Err()
	}
	return nil, cancel, fmt.Errorf("%s: max number of retries exceeded", url)
}
