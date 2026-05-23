package crawler

import (
	"encoding/json"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
}

func CrawlAPIArticles(link, category string) ([]model.Article, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching articles by api: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch articles: status %d", resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %w", err)
	}

	var apiResp []APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v\n", err)
	}

	articles := make([]model.Article, 0, len(apiResp))
	for _, item := range apiResp {
		if item.Link != "" && item.Title.Rendered != "" {
			articles = append(articles, model.Article{
				Title:    item.Title.Rendered,
				URL:      item.Link,
				Category: category,
				Format:   "json",
			})
		}
	}

	return articles, nil
}
