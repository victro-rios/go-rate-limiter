package ratelimiter

func setConfigDefaultValues(config *Config) {
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
