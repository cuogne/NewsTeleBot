package test

import (
	"fmt"
	"hcmus-news-tele-bot/internal/service"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func TestSummarizeWithGemini(t *testing.T) {
	var link string = "https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=292&newsid=17281"
	content, err := service.GetContentFromURL(link)
	if err != nil {
		t.Fatalf("Failed to get content: %v", err)
	}

	data, err := service.SummarizeContentWithGemini(content, link)
	if err != nil {
		t.Fatalf("Failed to summarize content: %v", err)
	}

	fmt.Println("Summary:", data.Summary)
}
