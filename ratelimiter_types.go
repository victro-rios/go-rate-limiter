package ratelimiter

import (
	"context"
)

type StoreClient interface {
	Get(ctx context.Context, key string) (*int32, error)
	Set(ctx context.Context, key string, value int32) error
}

type RateLimitHeaders struct {
	RetryAfter            string
	X_RateLimit_Limit     string
	X_RateLimit_Remaining string
	X_RateLimit_Reset     string
}

type RateLimiter struct {
	cfg        Config
	nextRefill int32
}
