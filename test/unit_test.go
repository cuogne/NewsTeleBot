package test

import (
	"encoding/json"
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/crawler"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/service"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNormalizeArticleURL(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "trim spaces and upgrade http to https",
			raw:  " http://hcmus.edu.vn/news?id=1 ",
			want: "https://hcmus.edu.vn/news?id=1",
		},
		{
			name: "keep https url unchanged",
			raw:  "https://hcmus.edu.vn/news",
			want: "https://hcmus.edu.vn/news",
		},
		{
			name: "add https to scheme relative url",
			raw:  "//hcmus.edu.vn/news",
			want: "https://hcmus.edu.vn/news",
		},
		{
			name: "keep plain host as raw value",
			raw:  "hcmus.edu.vn/news",
			want: "hcmus.edu.vn/news",
		},
		{
			name: "return empty string",
			raw:  "   ",
			want: "",
		},
		{
			name: "return trimmed raw value when url cannot parse",
			raw:  " http://%41:80/ ",
			want: "http://%41:80/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.NormalizeArticleURL(tt.raw)
			if got != tt.want {
				t.Fatalf("NormalizeArticleURL(%q) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}
}

func TestCrawlAPIArticles(t *testing.T) {
	var sawUserAgent bool

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", r.Method)
		}
		sawUserAgent = r.Header.Get("User-Agent") != ""

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]crawler.APIResponse{
			newAPIResponse("Tin hop le 1", "https://hcmus.edu.vn/1"),
			newAPIResponse("", "https://hcmus.edu.vn/missing-title"),
			newAPIResponse("Missing link", ""),
			newAPIResponse("Tin hop le 2", "https://hcmus.edu.vn/2"),
		})
	}))
	defer server.Close()

	got, err := crawler.CrawlAPIArticles(server.URL, "hcmus")
	if err != nil {
		t.Fatalf("CrawlAPIArticles() error = %v", err)
	}

	want := []model.Article{
		{Title: "Tin hop le 1", URL: "https://hcmus.edu.vn/1", Category: "hcmus", Format: "json"},
		{Title: "Tin hop le 2", URL: "https://hcmus.edu.vn/2", Category: "hcmus", Format: "json"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CrawlAPIArticles() = %#v, want %#v", got, want)
	}
	if !sawUserAgent {
		t.Fatal("expected CrawlAPIArticles to send a User-Agent header")
	}
}

func TestCrawlAPIArticlesReturnsError(t *testing.T) {
	t.Run("invalid request url", func(t *testing.T) {
		articles, err := crawler.CrawlAPIArticles("://bad-url", "hcmus")
		if err == nil {
			t.Fatalf("CrawlAPIArticles() error = nil, articles = %#v", articles)
		}
		if !strings.Contains(err.Error(), "failed to create API request") {
			t.Fatalf("error = %q, want it to contain request creation failure", err.Error())
		}
	})

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		wantErrContain string
	}{
		{
			name: "non ok status",
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "server error", http.StatusInternalServerError)
			},
			wantErrContain: "status 500",
		},
		{
			name: "invalid json",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{invalid-json"))
			},
			wantErrContain: "Error parsing JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handler)
			defer server.Close()

			articles, err := crawler.CrawlAPIArticles(server.URL, "hcmus")
			if err == nil {
				t.Fatalf("CrawlAPIArticles() error = nil, articles = %#v", articles)
			}
			if !strings.Contains(err.Error(), tt.wantErrContain) {
				t.Fatalf("error = %q, want it to contain %q", err.Error(), tt.wantErrContain)
			}
		})
	}
}

func TestCrawlUnsupportedFormat(t *testing.T) {
	ch := make(chan model.ListArticles, 1)

	crawler.Crawl(config.Resource{
		URL:      "https://hcmus.edu.vn",
		Category: "hcmus",
		Format:   "pdf",
	}, ch)

	got := <-ch
	if got.Err == nil {
		t.Fatalf("Crawl() error = nil, result = %#v", got)
	}
	if !strings.Contains(got.Err.Error(), "unsupported format: pdf") {
		t.Fatalf("error = %q, want unsupported format error", got.Err.Error())
	}
	if got.Category != "" {
		t.Fatalf("category = %q, want empty category for unsupported format", got.Category)
	}
	if got.Articles != nil {
		t.Fatalf("articles = %#v, want nil for unsupported format", got.Articles)
	}
}

func newAPIResponse(title, link string) crawler.APIResponse {
	var response crawler.APIResponse
	response.Title.Rendered = title
	response.Link = link
	return response
}
