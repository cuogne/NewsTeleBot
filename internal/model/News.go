package model

type (
	News struct {
		Title    string `json:"title"`
		URL      string `json:"url"`
		Category string `json:"category"`
		Format   string `json:"format"`
	}

	ListNews struct {
		News     []News
		Category string
		Err      error
	}
)
