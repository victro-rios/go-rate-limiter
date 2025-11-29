// package ratelimiter provides a middleware to stop exceeding requests given in configuration. For more information read README.md
package ratelimiter

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

func New(config Config) *RateLimiter {
	setConfigDefaultValues(&config)
	return &RateLimiter{cfg: config, buckets: make(map[string]*atomic.Int32)}
}

func (rateLimiter *RateLimiter) startRefilling(key string) {
	ticker := time.NewTicker(time.Duration(rateLimiter.cfg.PeriodDurationInSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		remaining := rateLimiter.buckets[key].Load()
		// If the capacity will overflow then maximum burst will be set
		rateLimiter.buckets[key].Store(min(rateLimiter.cfg.MaximumBurst, remaining+rateLimiter.cfg.RefillRatePerPeriod))
		rateLimiter.logger("refilling tokens")
	}
}

func (rateLimiter RateLimiter) logger(message string) {
	verbosePrefix := "RateLimiter:Logger:: "
	if !rateLimiter.cfg.Verbose || message == "" {
		return
	}

	fmt.Println(verbosePrefix + message)
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	_, exists := rateLimiter.buckets[key]
	if !exists {
		rateLimiter.buckets[key] = &atomic.Int32{}
		rateLimiter.buckets[key].Store(rateLimiter.cfg.MaximumBurst)
		rateLimiter.startRefilling(key)
	}
	rateLimiter.logger(fmt.Sprintf("consuming token. %d tokens left", rateLimiter.buckets[key].Load()))

	if rateLimiter.buckets[key].Load() <= 0 {
		rateLimiter.logger("throwing error 429")
		return errors.New("too many requests")
	}

	rateLimiter.buckets[key].Add(-1)
	return nil
}
