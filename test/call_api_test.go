package test

import (
	"hcmus-news-tele-bot/internal/crawler"
	"testing"
)

func TestCrawl(t *testing.T) {
	link := "https://www.ctda.hcmus.edu.vn/wp-json/wp/v2/posts?per_page=10&_fields=title,link"
	articles, err := crawler.CrawlCTDAByAPI(link, "ctda")
	if err != nil {
		t.Errorf("Error occurred while crawling CTDA articles: %v", err)
	}
	if len(articles) == 0 {
		t.Error("No articles found")
	}

	for i, article := range articles {
		t.Logf("Article %d: Title: %s, URL: %s", i+1, article.Title, article.URL)
	}
}
