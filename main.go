// package ratelimiter provides a middleware to stop exceeding requests given in configuration. For more information read README.md
package ratelimiter

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

func New(config Config) *RateLimiter {
	return &RateLimiter{cfg: config, buckets: make(map[string]*atomic.Int32)}
}

func (rateLimiter *RateLimiter) startRefilling(key string) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		remaining := rateLimiter.buckets[key].Load()
		// If the capacity will overflow then maximum burst will be set
		rateLimiter.buckets[key].Store(min(rateLimiter.cfg.MaximumBurst, rateLimiter.cfg.RefillRatePerMinute+remaining))
	}
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	_, exists := rateLimiter.buckets[key]
	if !exists {
		rateLimiter.buckets[key] = &atomic.Int32{}
		rateLimiter.buckets[key].Store(rateLimiter.cfg.MaximumBurst)
		rateLimiter.startRefilling(key)
	}
	rateLimiter.buckets[key].Add(-1)
	fmt.Printf("Consuming... %d tokens left", rateLimiter.buckets[key].Load())
	if rateLimiter.buckets[key].Load() <= 0 {
		fmt.Printf("throwing error 429")
		return errors.New("too many requests")
	}
	return nil
}
