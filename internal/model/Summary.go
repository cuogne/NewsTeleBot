package model

type (
	SummaryJob struct {
		Article  Article
		Category string
	}

	SummaryResult struct {
		Article         Article
		Category        string
		Summary         string
		PromptToken     int
		CompletionToken int
	}
)
