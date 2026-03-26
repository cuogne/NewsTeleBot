package service

import (
	"net/http"
	"net/url"

	"github.com/go-shiori/go-readability"
)

func GetContentFromURL(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	articleContent, err := readability.FromReader(resp.Body, parsedURL)
	if err != nil {
		return "", err
	}

	return articleContent.TextContent, nil
}
