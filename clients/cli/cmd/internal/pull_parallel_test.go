package internal

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateGate_AllowsConcurrentReaders(t *testing.T) {
	var gate sync.RWMutex
	var concurrent int64
	var maxConcurrent int64

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Simulate the read-lock gate pattern from downloadWithRateGate
			gate.RLock()
			gate.RUnlock()

			c := atomic.AddInt64(&concurrent, 1)
			for {
				old := atomic.LoadInt64(&maxConcurrent)
				if c <= old || atomic.CompareAndSwapInt64(&maxConcurrent, old, c) {
					break
				}
			}

			time.Sleep(time.Millisecond)
			atomic.AddInt64(&concurrent, -1)
		}()
	}
	wg.Wait()

	// RLock is shared, so all goroutines should run concurrently
	if maxConcurrent < 2 {
		t.Errorf("expected concurrent execution with RLock gate, max concurrent was %d", maxConcurrent)
	}
}

func TestRateGate_WriteLockBlocksAllReaders(t *testing.T) {
	var gate sync.RWMutex
	ready := make(chan struct{}, 4)

	// Simulate a rate-limited worker holding the write lock
	gate.Lock()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ready <- struct{}{}
			gate.RLock()
			gate.RUnlock()
		}()
	}

	// Wait for all goroutines to start
	for i := 0; i < 4; i++ {
		<-ready
	}
	time.Sleep(10 * time.Millisecond)

	// Release the write lock (simulating rate limit wait done)
	gate.Unlock()
	wg.Wait()
}

func TestRateGate_TryLockPreventsDoubleWait(t *testing.T) {
	var gate sync.RWMutex
	var waitCount int64

	// Simulate two workers hitting rate limit simultaneously
	gate.Lock() // first worker takes the write lock

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Second worker tries TryLock, should fail
		if gate.TryLock() {
			atomic.AddInt64(&waitCount, 1)
			gate.Unlock()
		}
	}()

	time.Sleep(10 * time.Millisecond)
	atomic.AddInt64(&waitCount, 1) // first worker counts
	gate.Unlock()
	wg.Wait()

	// Only 1 worker should have done the wait (the first one)
	if atomic.LoadInt64(&waitCount) != 1 {
		t.Errorf("expected exactly 1 rate limit wait, got %d", waitCount)
	}
}

func TestBuildDownloadOpts_DefaultFileFormat(t *testing.T) {
	target := &Target{
		File:      "locales/<locale_name>.json",
		ProjectID: "proj1",
	}
	localeFile := &LocaleFile{
		FileFormat: "json",
		Tag:        "",
	}

	opts, err := target.buildDownloadOpts(localeFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.FileFormat.Value() != "json" {
		t.Errorf("expected file format 'json', got %q", opts.FileFormat.Value())
	}
}

func TestBuildDownloadOpts_TagHandling(t *testing.T) {
	target := &Target{
		File:      "locales/<locale_name>/<tag>.json",
		ProjectID: "proj1",
	}
	localeFile := &LocaleFile{
		FileFormat: "json",
		Tag:        "web",
	}

	opts, err := target.buildDownloadOpts(localeFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Tags.Value() != "web" {
		t.Errorf("expected tags 'web', got %q", opts.Tags.Value())
	}
	if opts.Tag.Value() != "" {
		t.Errorf("expected tag to be empty string, got %q", opts.Tag.Value())
	}
}
