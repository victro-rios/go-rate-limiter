package ratelimiter

import (
	"context"
	"sync/atomic"
)

type StoreClient interface {
	Get(ctx context.Context, key string) (*int32, error)
	Set(ctx context.Context, key string, value int32) error
}

type MemoryStoreClient struct {
	buckets map[string]*atomic.Int32
}

type RateLimitHeaders struct {
	RetryAfter            int32
	X_RateLimit_Limit     int32
	X_RateLimit_Remaining int32
	X_RateLimit_Reset     int32
}

type RateLimiter struct {
	cfg        Config
	nextRefill int32
}
