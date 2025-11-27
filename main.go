package ratelimiter

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type RateLimiter struct {
	cfg     Config
	buckets map[string]*atomic.Int32
}

func New(config Config) *RateLimiter {
	return &RateLimiter{cfg: config, buckets: make(map[string]*atomic.Int32)}
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	_, exists := rateLimiter.buckets[key]
	if !exists {
		rateLimiter.buckets[key] = &atomic.Int32{}
		rateLimiter.buckets[key].Store(10)
	}
	rateLimiter.buckets[key].Add(1)
	fmt.Printf("Consuming... %d tokens left", rateLimiter.buckets[key].Load())
	if rateLimiter.buckets[key].Load() <= 0 {
		fmt.Printf("Throwing error 429")
		return errors.New("Too many requests")
	}
	return nil
}
