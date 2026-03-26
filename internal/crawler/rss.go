package crawler

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"time"

	"github.com/mmcdole/gofeed"
)

func CrawlRSSNews(link, category string) ([]model.News, error) {
	parser := gofeed.NewParser()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	feed, err := parser.ParseURLWithContext(link, ctx)

	if err != nil {
		return nil, err
	}

	var news []model.News

	for _, item := range feed.Items {
		if item.Title != "" && item.Link != "" {
			news = append(news, model.News{
				Category: category,
				Title:    item.Title,
				URL:      item.Link,
				Format:   "rss",
			})
		}
	}

	if len(news) == 0 {
		return nil, fmt.Errorf("no news element found for %s", category)
	}

	sizeNews := min(len(news), 10)
	news = news[:sizeNews]

	return news, nil
}
