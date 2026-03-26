package crawler

import (
	"fmt"
	"hcmus-news-tele-bot/internal/model"

	"github.com/gocolly/colly"
)

func CrawlHTMLNews(link, category string) ([]model.Article, error) {
	var articles []model.Article
	var crawlError error

	c := colly.NewCollector(
		colly.AllowedDomains("hcmus.edu.vn"),
	)

	c.OnHTML(".content.entry .cmsmasters_archive_item_title.entry-title a",
		func(e *colly.HTMLElement) {
			title := e.Text
			url := e.Attr("href")

			if title != "" && url != "" {
				articles = append(articles, model.Article{
					Category: category,
					Title:    title,
					URL:      url,
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

	if len(articles) == 0 {
		return nil, fmt.Errorf("no news element found for %s", category)
	}

	sizeNews := min(len(articles), 10)
	articles = articles[:sizeNews]

	// for i, j := 0, len(articles)-1; i < j; i, j = i+1, j-1 {
	// 	articles[i], articles[j] = articles[j], articles[i]
	// }

	return articles, nil
}
