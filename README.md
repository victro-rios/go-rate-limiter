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
    // could be GET or any other method
	mux.HandleFunc("POST /sampleEndpoint", func(w http.ResponseWriter, r *http.Request) {
		// taking request.Host as a key
        err := rateLimiter.Consume(r.Host, 1)
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
			return
		}
        // normal behavior 
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Thanks for adding host: %s", r.Host)))
	})
```


