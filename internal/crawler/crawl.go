package crawler

import (
	"fmt"
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/model"
	"strings"
)

// https://dev.to/jones_charles_ad50858dbc0/building-a-high-concurrency-web-crawler-in-go-a-practical-guide-i3a
func Crawl(feed config.Resource, ch chan<- model.ListArticles) {
	switch strings.ToLower(feed.Format) {

	// rss
	case "rss":
		listArticles, err := CrawlRSSArticles(feed.URL, feed.Category)

		ch <- model.ListArticles{
			Articles: listArticles,
			Category: feed.Category,
			Err:      err,
		}

	// html
	case "html":
		listArticles, err := CrawlHTMLArticles(feed.URL, feed.Category)

		ch <- model.ListArticles{
			Articles: listArticles,
			Category: feed.Category,
			Err:      err,
		}

	// api(json)
	case "api":
		listArticles, err := CrawlCTDAByAPI(feed.URL, feed.Category)

		ch <- model.ListArticles{
			Articles: listArticles,
			Category: feed.Category,
			Err:      err,
		}

	// default
	default:
		ch <- model.ListArticles{
			Articles: nil,
			Category: "",
			Err:      fmt.Errorf("unsupported format: %s", feed.Format),
		}
	}

}
