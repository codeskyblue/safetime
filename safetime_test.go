package safetime

import (
	"sync"
	"testing"
	"time"
)

func TestSafeTimer(tt *testing.T) {
	deadtime := time.Now().Add(2 * time.Second)
	t := NewTimer(100 * time.Hour)
	wg := sync.WaitGroup{}
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				t.Reset(10 * time.Hour)
				if time.Now().After(deadtime) {
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
