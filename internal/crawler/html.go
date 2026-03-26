package crawler

import (
	"fmt"
	"hcmus-news-tele-bot/internal/model"

	"github.com/gocolly/colly"
)

func CrawlHTMLNews(link, category string) ([]model.News, error) {
	var news []model.News
	var crawlError error

	c := colly.NewCollector(
		colly.AllowedDomains("hcmus.edu.vn"),
	)

	c.OnHTML(".content.entry .cmsmasters_archive_item_title.entry-title a",
		func(e *colly.HTMLElement) {
			title := e.Text
			url := e.Attr("href")

			if title != "" && url != "" {
				news = append(news, model.News{
					Category: category,
					Title:    title,
					URL:      url,
					Format:   "html",
				})
			}
		},
	)

	c.OnError(func(r *colly.Response, err error) {
		crawlError = fmt.Errorf("error crawling %s: %v", link, err)
	})

	err := c.Visit(link)
	if err != nil {
		return nil, fmt.Errorf("failed to visit %s: %v", link, err)
	}

	c.Wait()

	if crawlError != nil {
		return nil, crawlError
	}

	if len(news) == 0 {
		return nil, fmt.Errorf("no news element found for %s", category)
	}

	sizeNews := min(len(news), 10)
	news = news[:sizeNews]

	// for i, j := 0, len(news)-1; i < j; i, j = i+1, j-1 {
	// 	news[i], news[j] = news[j], news[i]
	// }

	return news, nil
}
