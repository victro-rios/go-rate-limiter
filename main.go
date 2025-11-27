package ratelimiter

import (
	"errors"
	"fmt"
)

type RateLimiter struct {
	cfg     Config
	buckets map[string]int
}

func New(config Config) *RateLimiter {
	return &RateLimiter{cfg: config}
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	_, exists := rateLimiter.buckets[key]
	if !exists {
		rateLimiter.buckets[key] = rateLimiter.cfg.MaximumBurst
	}
	rateLimiter.buckets[key] -= 1
	fmt.Printf("Consuming... %d tokens left", rateLimiter.buckets[key])
	if rateLimiter.buckets[key] <= 0 {
		fmt.Printf("Throwing error 429")
		return errors.New("Too many requests")
	}
	return nil
}
