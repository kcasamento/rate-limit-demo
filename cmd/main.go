package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kcasamento/rate-limit-demo/internal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	rateLimiter := internal.NewRateLimiter(3, time.Duration(10)*time.Second, ctx, true)

	go lambda(1, rateLimiter, ctx)
	go lambda(2, rateLimiter, ctx)
	go lambda(3, rateLimiter, ctx)
	go lambda(4, rateLimiter, ctx)
	go lambda(5, rateLimiter, ctx)
	go lambda(6, rateLimiter, ctx)
	go lambda(7, rateLimiter, ctx)
	go lambda(1, rateLimiter, ctx)
	go lambda(2, rateLimiter, ctx)
	go lambda(3, rateLimiter, ctx)
	go lambda(4, rateLimiter, ctx)
	go lambda(5, rateLimiter, ctx)
	go lambda(6, rateLimiter, ctx)
	go lambda(7, rateLimiter, ctx)

	oscall := <-c
	fmt.Printf("system call: %+v", oscall)
	cancel()
}

func lambda(taskId int, r *internal.RateLimiter, ctx context.Context) {
	for {
		select {
		case <-time.After(time.Duration(500) * time.Millisecond):
			fmt.Printf("Sending new request from lambda: %d\n", taskId)
			err := r.MakeRequest()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case <-ctx.Done():
			fmt.Println("Stopping lambda")
		}
	}
}
