package internal

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	items      map[string]int
	ctx        context.Context
	TTL        time.Duration
	threadSafe bool
	sync.RWMutex
}

func NewCache(ttl time.Duration, ctx context.Context, threadSafe bool) *Cache {
	c := &Cache{
		items:      make(map[string]int),
		ctx:        ctx,
		TTL:        ttl,
		threadSafe: threadSafe,
	}

	c.items["global"] = 0

	go c.RunItemTimer()
	return c
}

func (c *Cache) RunItemTimer() {
	for {
		select {
		case <-time.After(c.TTL):
			fmt.Println("Resetting cache")
			c.ResetItems()
		case <-c.ctx.Done():
			fmt.Println("Cancel received")
		}
	}
}

func (c *Cache) ResetItems() {
	if c.threadSafe {
		c.Lock()
		defer c.Unlock()
	}

	for key, value := range c.items {
		fmt.Printf("Resetting %s from %d to 0\n", key, value)
		c.items[key] = 0
	}
}

func (c *Cache) Inc(key string) {
	c.items[key]++
}

func (c *Cache) Get(key string) int {
	val := c.items[key]
	return val
}
