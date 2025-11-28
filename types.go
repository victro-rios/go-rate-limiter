package ratelimiter

import "sync/atomic"

type Config struct {
	MaximumBurst        int32
	TokensPerBucket     int32
	RefillRatePerMinute int32
}

type RateLimiter struct {
	cfg     Config
	buckets map[string]*atomic.Int32
}
