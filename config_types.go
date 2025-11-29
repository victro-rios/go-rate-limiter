package ratelimiter

type Config struct {
	MaximumBurst            int32
	TokensPerBucket         int32
	RefillRatePerPeriod     int32
	PeriodDurationInSeconds int32
	Verbose                 bool
}
