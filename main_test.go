package ratelimiter

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNoConfig(t *testing.T) {
	rateLimiter := New(Config{})
	logAssert(t, assert.Equal, "Should set the Maximum burst to 100", 100, rateLimiter.cfg.MaximumBurst)
	logAssert(t, assert.Equal, "Should set the Refill rate per period to 10", 10, rateLimiter.cfg.RefillRatePerPeriod)
	logAssert(t, assert.Equal, "Should set the period duration in seconds to 60", 60, rateLimiter.cfg.PeriodDurationInSeconds)
	sameType := assert.IsType(t, &MemoryStoreClient{}, rateLimiter.cfg.StoreClient)
	logAssert(t, assert.IsType, "Should set the default store to memory", sameType, true)
}

func TestLogger(t *testing.T) {
	cases := []struct {
		name    string
		message string
		prefix  string
	}{
		{name: "Should log in case a message is given with the prefix assigned (message: message)", prefix: "RateLimiter:Logger:", message: "message"},
		{name: "Should log in case a message is given with the prefix assigned (message: message 2)", prefix: "RateLimiter:Logger:", message: "message 2"},
		{name: "Should not log if message is not provided", prefix: "RateLimiter:Logger:"},
	}
	rateLimiter := New(Config{
		Verbose: true,
	})
	for _, tc := range cases {
		terminalOutput := captureTerminalOutput(func() {
			rateLimiter.logger(tc.message)
		})
		var expected string
		if tc.message != "" {
			expected = tc.prefix + tc.message + "\n"
		} else {
			expected = ""
		}
		got := string(terminalOutput)
		logAssert(t, assert.Equal, tc.name, expected, got)
	}
}

func logAssert[T any](t *testing.T, assertFunc func(t assert.TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool, testName string, expected T, actual T) {
	result := assertFunc(t, expected, actual)

	if !result {
		t.Errorf("❌ %s: Expected %v, Got %v", testName, expected, actual)
	} else {
		t.Logf("✔️ %s", testName)
	}
}

func captureTerminalOutput(f func()) string {
	// Capture Stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	// Stop capturing
	w.Close()
	os.Stdout = old

	terminalOutput, _ := io.ReadAll(r)
	return string(terminalOutput)
}
