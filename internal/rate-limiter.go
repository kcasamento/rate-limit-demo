package internal

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type RateLimiter struct {
	cache     *Cache
	rateLimit int
}

func NewRateLimiter(rateLimit int, rateInterval time.Duration, ctx context.Context) *RateLimiter {
	cache := NewCache(rateInterval, ctx)
	return &RateLimiter{
		cache:     cache,
		rateLimit: rateLimit,
	}
}

func (r *RateLimiter) MakeRequest() error {
	r.cache.Lock()
	defer r.cache.Unlock()

	currVal := r.cache.Get("global")
	fmt.Printf("Current rate limit: %d\n", currVal)

	if currVal <= r.rateLimit {
		r.cache.Inc("global")
		fmt.Printf("Under rate limit...making request")
		return nil
	}

	return errors.New("you have been throttled")
}
