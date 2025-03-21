package concurrency

import ()

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// returns a map of each URL checked to a boolean value: true for a good response; false for a bad response
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
