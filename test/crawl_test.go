package test

import (
	"testing"

	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/service"
)

func TestGetArticles(t *testing.T) {
	articles, err := service.GetArticles()
	if err != nil {
		t.Fatalf("Đã xảy ra lỗi khi lấy bài báo: %v", err)
	}

	if len(articles) == 10*len(config.Feeds) {
		t.Logf("PASS CRAWL, get %d News", len(articles))
	} else {
		t.Logf("FAIL CRAWL, just got %d News", len(articles))
	}
}
