package model

type (
	Article struct {
		Title    string `json:"title"`
		URL      string `json:"url"`
		Category string `json:"category"`
		Format   string `json:"format"`
	}

	ListArticles struct {
		Articles []Article
		Category string
		Err      error
	}
)
