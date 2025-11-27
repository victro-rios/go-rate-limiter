package ratelimiter

type Config struct {
	MaximumBurst atomic.Uint32
}
