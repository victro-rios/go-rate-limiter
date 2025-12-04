# go-rate-limiter
**go-rate-limiter** is a middleware that provides an easy to use module for golang developers. Allowing them to quickly integrate a
rate limiter in a few simple steps. 

## Installation
`go get github.com/victro-rios/go-rate-limiter`

`go mod tidy`

## Usage
```golang
	rateLimiterConfig := ratelimiter.Config{
		MaximumBurst:            10,
		RefillRatePerPeriod:     5,
		PeriodDurationInSeconds: 60,
		Verbose:                 true,
	}
	rateLimiter := ratelimiter.New(rateLimiterConfig)

	mux := http.NewServeMux()
    // could be GET/POST or any other method
	mux.HandleFunc("POST /sampleEndpoint", func(w http.ResponseWriter, r *http.Request) {
		// taking request.RemoteAddr as a key
        err := rateLimiter.Consume(r.RemoteAddr, 1)
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
			return
		}
        // normal behavior 
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
```

## Setting headers

```golang
	rateLimiterConfig := ratelimiter.Config{}
	rateLimiter := ratelimiter.New(rateLimiterConfig)

	mux := http.NewServeMux()
    // could be GET/POST or any other method
	mux.HandleFunc("POST /sampleEndpoint", func(w http.ResponseWriter, r *http.Request) {
		// taking request.RemoteAddr as a key
        err := rateLimiter.Consume(r.RemoteAddr, 1)
		if err != nil {
			w.Header().Set("Retry-After", err.Headers.RetryAfter)
			w.Header().Set("X-RateLimit-Limit", err.Headers.X_RateLimit_Limit)
			w.Header().Set("X-RateLimit-Remaining", err.Headers.X_RateLimit_Remaining)
			w.Header().Set("X-RateLimit-Reset", err.Headers.X_RateLimit_Reset)
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
			return
		}
        // normal behavior 
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
```
