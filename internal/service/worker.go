package service

import (
	"hcmus-news-tele-bot/internal/model"
	"strings"
)

func Worker(id int, jobs <-chan model.SummaryJob, results chan<- model.SummaryResult) {
	for job := range jobs {
		category := strings.ToLower(job.Category)
		if category == "lichthi" || category == "thongbao" {
			results <- model.SummaryResult{
				Article:  job.Article,
				Category: job.Category,
				Summary:  job.Article.URL,
			}
			continue
		}

		content, err := GetContentFromURL(job.Article.URL)
		if err != nil {
			results <- model.SummaryResult{
				Article:  job.Article,
				Category: job.Category,
				Summary:  "",
			}
			continue
		}

		summary, err := SummarizeContentWithGemini(content, job.Article.URL)
		if err != nil {
			results <- model.SummaryResult{
				Article:  job.Article,
				Category: job.Category,
				Summary:  "",
			}
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
