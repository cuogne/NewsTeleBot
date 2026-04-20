package service

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-shiori/go-readability"
)

func GetContentFromURL(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(parsedURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch URL: status %d", resp.StatusCode)
	}

	articleContent, err := readability.FromReader(resp.Body, parsedURL)
	if err != nil {
		return "", err
	}

	return articleContent.TextContent, nil
}
