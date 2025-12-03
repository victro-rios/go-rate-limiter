package ratelimiter

import "sync/atomic"

func setConfigDefaultValues(config *Config) {
	if config.StoreClient == nil {
		config.StoreClient = &MemoryStoreClient{
			buckets: make(map[string]*atomic.Int32),
		}
	}
	if config.MaximumBurst == 0 {
		config.MaximumBurst = 100
	}

	if config.RefillRatePerPeriod == 0 {
		config.RefillRatePerPeriod = 10
	}

	if config.PeriodDurationInSeconds == 0 {
		config.PeriodDurationInSeconds = 60
	}
}
