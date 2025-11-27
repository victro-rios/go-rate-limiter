package ratelimiter

type Config struct {
	MaximumBurst        int
	TokensPerBucket     int
	RefillRatePerSecond int
}
