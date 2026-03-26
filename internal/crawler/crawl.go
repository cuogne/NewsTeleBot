package crawler

import (
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/model"
	"strings"
	"sync"
)

// https://dev.to/jones_charles_ad50858dbc0/building-a-high-concurrency-web-crawler-in-go-a-practical-guide-i3a
func Crawl(feed config.Resource, wg *sync.WaitGroup, ch chan<- model.ListNews) {
	defer wg.Done()

	switch strings.ToLower(feed.Format) {
	case "rss":
		listNews, err := CrawlRSSNews(feed.URL, feed.Category)

		ch <- model.ListNews{
			News:     listNews,
			Category: feed.Category,
			Err:      err,
		}
	case "html":
		listNews, err := CrawlHTMLNews(feed.URL, feed.Category)

		ch <- model.ListNews{
			News:     listNews,
			Category: feed.Category,
			Err:      err,
		}
	}
}
