package service

import "sync"

// maxFetchAttempts limits how many cron cycles retry a failed content fetch
// before falling back to sending the article with an empty summary.
// The attempt counter is in-memory only, so it resets on restart.
const maxFetchAttempts = 3

var (
	fetchMu       sync.Mutex
	fetchAttempts = make(map[string]int)
)

// shouldRetryFetch records a failed fetch for the url and reports
// whether it should be retried on the next cron cycle.
func shouldRetryFetch(url string) bool {
	fetchMu.Lock()
	defer fetchMu.Unlock()

	fetchAttempts[url]++
	return fetchAttempts[url] < maxFetchAttempts
}

// clearFetchAttempt removes the retry counter after the url is processed.
func clearFetchAttempt(url string) {
	fetchMu.Lock()
	defer fetchMu.Unlock()

	delete(fetchAttempts, url)
}
