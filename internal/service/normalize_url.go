package service

import (
	"net/url"
	"strings"
)

func NormalizeArticleURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	if strings.EqualFold(parsedURL.Scheme, "http") {
		parsedURL.Scheme = "https"
		return parsedURL.String()
	}

	if parsedURL.Scheme == "" && parsedURL.Host != "" {
		parsedURL.Scheme = "https"
		return parsedURL.String()
	}

	return rawURL
}
