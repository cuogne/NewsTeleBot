package service

import (
	"fmt"
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/crawler"
	"hcmus-news-tele-bot/internal/model"
	"log"
	"sync"
)

func GetArticles() ([]model.Article, error) {
	// urls := make([]string, len(config.Feeds))

	// for i, fd := range config.Feeds {
	// 	urls[i] = fd.URL
	// }

	var wg sync.WaitGroup
	ch := make(chan model.ListArticles, len(config.Feeds))
	sem := make(chan struct{}, 10) // max concurrency 10

	for _, fd := range config.Feeds {
		wg.Add(1)

		go func(feed config.Resource) {
			defer wg.Done()

			// use semaphore
			sem <- struct{}{}
			defer func() { <-sem }()

			crawler.Crawl(feed, ch)
		}(fd)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var articles = make([]model.Article, 0, len(config.Feeds)*10)
	failedFeeds := 0

	for listArticles := range ch {
		if listArticles.Err != nil {
			log.Printf("err: %v", listArticles.Err)
			failedFeeds++
			continue // skip error feed
		}
		for i := range listArticles.Articles {
			listArticles.Articles[i].Category = listArticles.Category
		}
		articles = append(articles, listArticles.Articles...)
	}

	if len(articles) == 0 && failedFeeds > 0 {
		return nil, fmt.Errorf("crawl failed for all feeds (%d/%d)", failedFeeds, len(config.Feeds))
	}

	return articles, nil
}
