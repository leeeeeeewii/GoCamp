package rollingWindow

import (
	"sync"
	"time"
)

type RollingWindow struct {
		lock          sync.RWMutex
		size          int
		win           *window
		interval      time.Duration
		offset        int
		ignoreCurrent bool
		lastTime      time.Duration
	}

// Add adds value to current bucket.
func (rw *RollingWindow) Add(v float64) {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.updateOffset()
	rw.win.add(rw.offset, v)
}

// Reduce runs function fn on all buckets.
func (rw *RollingWindow) Reduce(fn func(b *Bucket)) {
	rw.lock.RLock()
	defer rw.lock.RUnlock()

	var diff int
	span := rw.span()
  
	if span == 0 && rw.ignoreCurrent {
		diff = rw.size - 1
	} else {
		diff = rw.size - span
	}
	if diff > 0 {
		offset := (rw.offset + span + 1) % rw.size
		rw.win.reduce(offset, diff, fn)
	}
}

// span shows the number of expired buckets
func (rw *RollingWindow) span() int {
  since_diff = time.Since(time.Now().AddDate(-1, -1, -1)) - rw.lastTime
  offset := int(since_diff / rw.interval)
	if 0 <= offset && offset < rw.size {
		return offset
	}

	return rw.size
}

// updateOffset calculates the offset
func (rw *RollingWindow) updateOffset() {
	span := rw.span()
	if span <= 0 {
		return
	}

	offset := rw.offset
	// reset expired buckets
	for i := 0; i < span; i++ {
		rw.win.resetBucket((offset + i + 1) % rw.size)
	}

	rw.offset = (offset + span) % rw.size
	now := time.Since(time.Now().AddDate(-1, -1, -1))
	// align to interval time boundary
	rw.lastTime = now - (now-rw.lastTime)%rw.interval
}
