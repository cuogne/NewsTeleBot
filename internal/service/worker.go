package service

import (
	"hcmus-news-tele-bot/internal/model"
	"strings"
)

func Worker(id int, jobs <-chan model.SummaryJob, results chan<- model.SummaryResult) {
	for job := range jobs {
		category := strings.ToLower(job.Category)

		// for "lichthi" and "thongbao", do not summarize (it's so short)
		if category == "lichthi" || category == "thongbao" {
			results <- model.SummaryResult{
				Article:         job.Article,
				Category:        job.Category,
				Summary:         job.Article.URL,
				PromptToken:     0,
				CompletionToken: 0,
			}
			continue
		}

		content, err := GetContentFromURL(job.Article.URL)
		if err != nil {
			results <- model.SummaryResult{
				Article:         job.Article,
				Category:        job.Category,
				Summary:         "",
				PromptToken:     0,
				CompletionToken: 0,
			}
			continue
		}

		summary, err := SummarizeContentWithGemini(content)
		if err != nil {
			results <- model.SummaryResult{
				Article:         job.Article,
				Category:        job.Category,
				Summary:         "",
				PromptToken:     0,
				CompletionToken: 0,
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
