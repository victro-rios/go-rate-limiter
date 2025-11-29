package ratelimiter

import "sync/atomic"

type RateLimiter struct {
	cfg     Config
	buckets map[string]*atomic.Int32
}
