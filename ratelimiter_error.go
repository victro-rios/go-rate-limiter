package ratelimiter

type RateLimiterError struct {
	Msg     string
	Code    int32
	Headers RateLimitHeaders
}

func (rateLimiterError RateLimiterError) Error() string {
	return rateLimiterError.Msg
}
