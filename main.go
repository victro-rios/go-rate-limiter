// package ratelimiter provides a middleware to stop exceeding requests given in configuration. For more information read README.md
package ratelimiter

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func New(config Config) *RateLimiter {
	setConfigDefaultValues(&config)
	return &RateLimiter{cfg: config}
}

func (rateLimiter *RateLimiter) startRefilling(key string) error {
	ticker := time.NewTicker(time.Duration(rateLimiter.cfg.PeriodDurationInSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		remaining, err := rateLimiter.cfg.StoreClient.Get(context.Background(), key)
		if err != nil {
			return errors.New("error getting key from store client")
		}
		// If the capacity will overflow then maximum burst will be set
		rateLimiter.cfg.StoreClient.Set(context.Background(), key, min(rateLimiter.cfg.MaximumBurst, *remaining+rateLimiter.cfg.RefillRatePerPeriod))
		rateLimiter.logger("refilling tokens")
	}
	return nil
}

func (rateLimiter RateLimiter) logger(message string) {
	verbosePrefix := "RateLimiter:Logger:"
	if !rateLimiter.cfg.Verbose || message == "" {
		return
	}

	fmt.Println(verbosePrefix + message)
}

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) error {
	keyValue, err := rateLimiter.cfg.StoreClient.Get(context.Background(), key)
	if err != nil {
		return errors.New("error while consuming the key from store client")
	}

	if keyValue == nil {
		rateLimiter.cfg.StoreClient.Set(context.Background(), key, rateLimiter.cfg.MaximumBurst)
		rateLimiter.startRefilling(key)
		keyValue, _ = rateLimiter.cfg.StoreClient.Get(context.Background(), key)
	}
	rateLimiter.logger(fmt.Sprintf("consuming token. %d tokens left", *keyValue))

	if *keyValue <= 0 {
		rateLimiter.logger("throwing error 429")
		return errors.New("too many requests")
	}
	// TODO: Check concurrency
	rateLimiter.cfg.StoreClient.Set(context.Background(), key, int32(*keyValue-1))
	return nil
}
