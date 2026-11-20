package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimitTransport(t *testing.T) {
	now := time.Unix(1000, 0)
	prevSleep, prevNow := rateLimitSleep, rateLimitNow
	t.Cleanup(func() { rateLimitSleep, rateLimitNow = prevSleep, prevNow })
	rateLimitNow = func() time.Time { return now }

	cases := []struct {
		name      string
		reason    string
		reset     int64
		rejects   int32
		wantCalls int32
		wantDelay time.Duration
		wantCode  int
	}{
		{"concurrency", "global-concurrency", 0, 2, 3, concurrencyDelay, 200},
		{"rate waits for reset", "global-rate", 1010, 1, 2, 11 * time.Second, 200},
		{"reset in the past", "global-rate", 900, 1, 2, concurrencyDelay, 200},
		{"gives up", "global-concurrency", 0, 100, maxRateLimitRetries + 1, concurrencyDelay, 429},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var calls int32
			var slept []time.Duration
			rateLimitSleep = func(d time.Duration) { slept = append(slept, d) }
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				if string(body) != "payload" {
					t.Errorf("body = %q, want payload", body)
				}
				if atomic.AddInt32(&calls, 1) <= c.rejects {
					w.Header().Set("X-Rate-Limit-Reason", c.reason)
					w.Header().Set("X-Rate-Limit-Remaining", "0")
					w.Header().Set("X-Rate-Limit-Reset", strconv.FormatInt(c.reset, 10))
					w.WriteHeader(http.StatusTooManyRequests)
					return
				}
			}))
			defer srv.Close()

			client := &http.Client{Transport: rateLimitTransport{http.DefaultTransport}}
			resp, err := client.Post(srv.URL, "text/plain", strings.NewReader("payload"))
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != c.wantCode || calls != c.wantCalls {
				t.Errorf("status=%d calls=%d, want %d/%d", resp.StatusCode, calls, c.wantCode, c.wantCalls)
			}
			if len(slept) == 0 || slept[0] != c.wantDelay {
				t.Errorf("slept %v, want first delay %s", slept, c.wantDelay)
			}
		})
	}
}
