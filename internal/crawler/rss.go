package crawler

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"time"

	"github.com/mmcdole/gofeed"
)

func CrawlRSSArticles(link, category string) ([]model.Article, error) {
	parser := gofeed.NewParser()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	feed, err := parser.ParseURLWithContext(link, ctx)

	if err != nil {
		return nil, err
	}

	var articles []model.Article

	for _, item := range feed.Items {
		if item.Title != "" && item.Link != "" {
			articles = append(articles, model.Article{
				Category: category,
				Title:    item.Title,
				URL:      item.Link,
				Format:   "rss",
			})
		}
	}

	if len(articles) == 0 {
		return nil, fmt.Errorf("no articles element found for %s", category)
	}

	sizeArticles := min(len(articles), 10)
	articles = articles[:sizeArticles]

	return articles, nil
}
