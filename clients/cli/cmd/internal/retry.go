package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/phrase/phrase-cli/cmd/internal/print"
)

const (
	maxDownloadAttempts = 3
	maxRateLimitRetries = 20
	concurrencyDelay    = 250 * time.Millisecond
)

var (
	retrySleep   = sleepContext
	rateLimitNow = time.Now
)

// sleepContext waits for d or until ctx is done, whichever comes first.
func sleepContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// retryTransient calls fn up to maxDownloadAttempts times for as long as it
// fails with an error it reports as transient.
func retryTransient(fn func() (transient bool, err error)) error {
	for attempt := 1; ; attempt++ {
		transient, err := fn()
		if err == nil || !transient || attempt >= maxDownloadAttempts {
			return err
		}
		delay := time.Second << (attempt - 1) // 1s, 2s
		debugFprintln(fmt.Sprintf("Download attempt %d/%d failed: %s, retrying in %s", attempt, maxDownloadAttempts, err, delay))
		retrySleep(context.Background(), delay)
	}
}

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

		if resp.Header.Get("X-Rate-Limit-Reason") == "global-rate" {
			print.Warn("Rate limit exceeded. Retrying in %s", delay.Round(time.Second))
		} else {
			debugFprintln(fmt.Sprintf("Concurrency limit exceeded, retrying in %s", delay))
		}
		if err := retrySleep(req.Context(), delay); err != nil {
			if retry.Body != nil {
				retry.Body.Close()
			}
			return nil, err
		}
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
