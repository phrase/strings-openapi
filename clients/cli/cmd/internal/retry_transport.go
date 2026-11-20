package internal

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	maxRateLimitRetries = 20
	concurrencyDelay    = 250 * time.Millisecond
)

var (
	rateLimitSleep = time.Sleep
	rateLimitNow   = time.Now
)

// rateLimitTransport retries requests rejected with 429, either because of
// the global concurrency limit (retried after a short delay) or the global
// rate limit (retried once the window resets).
type rateLimitTransport struct {
	next http.RoundTripper
}

func (t rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 1; ; attempt++ {
		resp, err := t.next.RoundTrip(req)
		if err != nil || resp.StatusCode != http.StatusTooManyRequests || attempt > maxRateLimitRetries {
			return resp, err
		}
		delay := rateLimitDelay(resp)

		// the body is needed to replay the request
		retry := req.Clone(req.Context())
		if req.Body != nil && req.Body != http.NoBody {
			if req.GetBody == nil {
				return resp, nil
			}
			if retry.Body, err = req.GetBody(); err != nil {
				return resp, nil
			}
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		debugFprintln(fmt.Sprintf("Rate limited (%s), retrying in %s", resp.Header.Get("X-Rate-Limit-Reason"), delay))
		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		default:
		}
		rateLimitSleep(delay)
		req = retry
	}
}

func rateLimitDelay(resp *http.Response) time.Duration {
	if resp.Header.Get("X-Rate-Limit-Reason") != "global-rate" {
		return concurrencyDelay
	}
	// X-Rate-Limit-Remaining is 0 here; X-Rate-Limit-Reset is the unix time the window ends.
	reset, err := strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Reset"), 10, 64)
	if err != nil {
		return time.Second
	}
	return max(time.Unix(reset, 0).Sub(rateLimitNow())+time.Second, concurrencyDelay)
}
