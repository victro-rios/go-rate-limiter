package ratelimiter

type RateLimiterError struct {
	Msg     string
	Code    int
	Headers RateLimitHeaders
}

func (rateLimiterError RateLimiterError) Error() string {
	return rateLimiterError.Msg
}
