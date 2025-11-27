package ratelimiter

import (
	"fmt"
)

type RateLimiter struct {
	cfg Config
}

func New(config Config) *RateLimiter {
	return &RateLimiter{cfg: config}
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	fmt.Printf("Consuming")
	return nil
}
