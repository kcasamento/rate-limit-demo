# rate-limit-demo

## Cache
Mimics a fake "redis-like" cache that will implement a ttl. The ttl will trigger a goroutine to run and reset all cache items back to 0.  Run this demo with mutex locks on and off to see how this can affect the application.

## Rate Limiter
Sets up a fake "rate limiter" to simulate making http calls to a throttled resource.  At the time of writing this the default values were to allow 3 requests every 10 seconds.  This can be changed to any values you like.
Running the demo with the mutex locks on an off will showcase what happens in a rate limiting example.  When a lot of resources try to access the rate limit value without locking, it is very easy to go over the values set.
With locking, we can see that the rate limit will never go above the set maximum and will be properly reset on the interval specificied by ttl.

## Main
Setups the demo to run by creating a RateLimer object and simulating many lambda functions running on their own goroutine:
```
make run
```


