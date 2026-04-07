package service

import (
	"hcmus-news-tele-bot/internal/model"
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
			continue
		}

		summary, err := SummarizeContentWithGemini(content)
		if err != nil {
			results <- emptyResult()
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
