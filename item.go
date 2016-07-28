package ttlcache

import (
	"sync"
	"time"
)

// Item represents a record in the cache map
type Item struct {
	sync.RWMutex
	data    []string
	expires *time.Time
}

func (item *Item) touch(duration time.Duration) {
	// no-op
}

func (item *Item) expired() bool {
	var value bool
	item.RLock()
	if item.expires == nil {
		value = true
	} else {
		value = item.expires.Before(time.Now())
	}
	item.RUnlock()
	return value
}
