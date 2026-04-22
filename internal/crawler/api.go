package crawler

import (
	"encoding/json"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"io"
	"net/http"
	"time"
)

type ctdaAPIResponse struct {
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
}

func CrawlCTDAByAPI(link, category string) ([]model.Article, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", link, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching CTDA articles: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch CTDA articles: status %d", resp.StatusCode)
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)

	var apiResp []ctdaAPIResponse
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
