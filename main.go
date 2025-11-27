package ratelimiter

import (
	"sync/atomic"
	"fmt"
)

type RateLimiter struct {
	cfg: ConfConfig
}

func New(config Config) {
	return RateLimiter{cfg: config}
}

func Consume(key string, tokensToConsume uint8) error {
	fmt.Printf("Consuming");
	return nil
}
