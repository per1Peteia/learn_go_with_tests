package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureTime(a)
	bDuration := measureTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
