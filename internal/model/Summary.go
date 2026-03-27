package model

type (
	SummaryJob struct {
		Article  News
		Category string
	}

	SummaryResult struct {
		Article         News
		Category        string
		Summary         string
		PromptToken     int
		CompletionToken int
	}
)
