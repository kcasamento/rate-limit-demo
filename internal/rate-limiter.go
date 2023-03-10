package internal

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type RateLimiter struct {
	cache      *Cache
	rateLimit  int
	threadSafe bool
}

func NewRateLimiter(rateLimit int, rateInterval time.Duration, ctx context.Context, threadSafe bool) *RateLimiter {
	cache := NewCache(rateInterval, ctx, threadSafe)
	return &RateLimiter{
		cache:      cache,
		rateLimit:  rateLimit,
		threadSafe: threadSafe,
	}
}

func (r *RateLimiter) MakeRequest() error {
	if r.threadSafe {
		r.cache.Lock()
		defer r.cache.Unlock()
	}

	currVal := r.cache.Get("global")
	fmt.Printf("Current rate limit: %d\n", currVal)

	if currVal <= r.rateLimit {
		r.cache.Inc("global")
		fmt.Printf("Under rate limit...making request")
		return nil
	}

	return errors.New("you have been throttled")
}
