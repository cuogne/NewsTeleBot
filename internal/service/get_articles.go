package service

import (
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/crawler"
	"hcmus-news-tele-bot/internal/model"
	"log"
	"sync"
)

func GetArticles() ([]model.News, error) {
	// urls := make([]string, len(config.Feeds))

	// for i, fd := range config.Feeds {
	// 	urls[i] = fd.URL
	// }

	var wg sync.WaitGroup
	ch := make(chan model.ListNews, len(config.Feeds))

	for _, fd := range config.Feeds {
		wg.Add(1)
		go crawler.Crawl(fd, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var articles []model.News

	for listNews := range ch {
		if listNews.Err != nil {
			log.Fatal(listNews.Err)
			continue // skip error feed
		}
		articles = append(articles, listNews.News...)
	}

	return articles, nil
}
