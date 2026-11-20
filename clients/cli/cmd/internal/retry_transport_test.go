package internal

import (
	"context"
	"errors"
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
			rateLimitSleep = func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil }
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

func TestRateLimitTransportCancelledDuringBackoff(t *testing.T) {
	prevSleep := rateLimitSleep
	t.Cleanup(func() { rateLimitSleep = prevSleep })
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("X-Rate-Limit-Reason", "global-concurrency")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	rateLimitSleep = func(ctx context.Context, d time.Duration) error {
		cancel()
		return sleepContext(ctx, time.Hour)
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, srv.URL, strings.NewReader("payload"))
	_, err := rateLimitTransport{http.DefaultTransport}.RoundTrip(req)
	if !errors.Is(err, context.Canceled) || calls != 1 {
		t.Errorf("err=%v calls=%d, want context.Canceled after 1 call", err, calls)
	}
}
