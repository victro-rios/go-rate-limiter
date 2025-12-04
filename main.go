// package ratelimiter provides a middleware to stop exceeding requests given in configuration. For more information read README.md
package ratelimiter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func New(config Config) *RateLimiter {
	setConfigDefaultValues(&config)
	return &RateLimiter{cfg: config}
}

func (rateLimiter *RateLimiter) startRefilling(key string) error {
	interval := time.Duration(rateLimiter.cfg.PeriodDurationInSeconds) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	rateLimiter.nextRefill = int32(time.Now().Add(interval).Unix())
	for range ticker.C {
		remaining, err := rateLimiter.cfg.StoreClient.Get(context.Background(), key)
		if err != nil {
			return errors.New("error getting key from store client")
		}
		// If the capacity will overflow then maximum burst will be set
		rateLimiter.cfg.StoreClient.Set(context.Background(), key, min(rateLimiter.cfg.MaximumBurst, *remaining+rateLimiter.cfg.RefillRatePerPeriod))
		rateLimiter.logger("refilling tokens")
		rateLimiter.nextRefill = int32(time.Now().Add(interval).Unix())
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

func (rateLimiter *RateLimiter) Consume(key string, tokensToConsume uint8) *RateLimiterError {
	keyValue, err := rateLimiter.cfg.StoreClient.Get(context.Background(), key)
	if err != nil {
		return &RateLimiterError{
			Msg:  "error while consuming the key from store client",
			Code: 500,
		}
	}
	if keyValue == nil {
		rateLimiter.cfg.StoreClient.Set(context.Background(), key, rateLimiter.cfg.MaximumBurst)
		go rateLimiter.startRefilling(key)
		keyValue, _ = rateLimiter.cfg.StoreClient.Get(context.Background(), key)
	}

	rateLimiter.logger(fmt.Sprintf("consuming token. %d tokens left", *keyValue))

	if *keyValue <= 0 {
		rateLimiter.logger("throwing error 429")
		return &RateLimiterError{
			Msg:  "too many requests",
			Code: http.StatusTooManyRequests,
			Headers: RateLimitHeaders{
				RetryAfter:            strconv.Itoa(int(rateLimiter.nextRefill - int32(time.Now().Unix()))),
				X_RateLimit_Limit:     strconv.Itoa(int(rateLimiter.cfg.MaximumBurst)),
				X_RateLimit_Remaining: "0",
				X_RateLimit_Reset:     strconv.Itoa(int(rateLimiter.nextRefill)),
			},
		}
	}
	// TODO: Check concurrency
	rateLimiter.cfg.StoreClient.Set(context.Background(), key, int32(*keyValue-1))
	return nil
}
