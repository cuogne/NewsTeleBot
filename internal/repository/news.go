package repository

import (
	"encoding/json"
	"fit-news-discord-bot/internal/supabase"
	"fmt"
)

type News struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	SentAt string `json:"sent_at"`
}

func GetAllNews(tableName string) ([]News, error) {
	var news []News

	res, _, err := supabase.Postgrest.
		From(tableName).
		Select("*", "", false).
		Execute()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &news)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling: %v", err)
	}

	return news, nil
}
