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
		Retryable       bool // true when fetch failed and should be retried next cycle
	}
)
