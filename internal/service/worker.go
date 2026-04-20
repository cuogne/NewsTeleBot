package service

import (
	"hcmus-news-tele-bot/internal/model"
	"log"
	"strings"
)

func Worker(id int, jobs <-chan model.SummaryJob, results chan<- model.SummaryResult) {
	for job := range jobs {
		category := strings.ToLower(job.Category)

		emptyResult := func() model.SummaryResult {
			return model.SummaryResult{
				Article:         job.Article,
				Category:        job.Category,
				Summary:         "",
				PromptToken:     0,
				CompletionToken: 0,
			}
		}

		// for "lichthipkt" and "thongbaopkt", do not summarize (it's so short)
		switch category {
		case "thongbaopkt", "lichthipkt":
			results <- emptyResult()
			continue
		}

		content, err := GetContentFromURL(job.Article.URL)
		if err != nil {
			results <- emptyResult()
			log.Printf("Error getting content while fetching URL %s: %v", job.Article.URL, err)
			continue
		}

		// content is empty -> no need to summarize, just return empty summary
		// ("", err) or ("", nil)
		if content == "" {
			results <- emptyResult()
			log.Printf("No content extracted from URL %s", job.Article.URL)
			continue
		}

		summary, err := SummarizeContentWithGemini(content)
		if err != nil {
			results <- emptyResult()
			log.Printf("Error summarizing content from URL %s: %v", job.Article.URL, err)
			continue
		}

		results <- model.SummaryResult{
			Article:         job.Article,
			Category:        job.Category,
			Summary:         summary.Summary,
			PromptToken:     summary.PromptToken,
			CompletionToken: summary.CompletionToken,
		}
	}
}
