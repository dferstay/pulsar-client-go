package internal

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	clockReadInterval = 1 * time.Millisecond
)

var (
	clockNow  int64
	clockOnce sync.Once
)

// StartClockSource runs a background go-routine that periodically samples the system clock
// NB: this call is idempotent
func StartClockSource() {
	clockOnce.Do(func() {
		ticker := time.NewTicker(clockReadInterval)
		go func() {
			for {
				select {
				case t := <-ticker.C:
					atomic.StoreInt64(&clockNow, t.UnixNano())
				}
			}
		}()
	})
}

// Now returns the current value of the sampled system clock in nanoseconds since Unix epoch
func Now() int64 {
	return atomic.LoadInt64(&clockNow)
}
