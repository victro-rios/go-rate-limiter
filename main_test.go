package ratelimiter

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		result := assert.Equal(t, expected, got)
		if !result {
			t.Errorf("❌ %s: Expected %s, Got %s", tc.name, expected, got)
		} else {
			t.Logf("✔️ %s", tc.name)
		}
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
