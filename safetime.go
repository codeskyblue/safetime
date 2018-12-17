package safetime

import (
	"sync"
	"time"
)

// SafeTime add thread-safe for time.Timer
type Timer struct {
	*time.Timer
	mu sync.Mutex
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		Timer: time.NewTimer(d),
	}
}

// Reset is thread-safe now
func (t *Timer) Reset(d time.Duration) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.Timer.Reset(d)
}
