package service

import (
	"context"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveToDB(dbPool *pgxpool.Pool, res model.SummaryResult) {
	if res.Article.URL == "" || res.Article.Title == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// save to db
	err := repository.SaveArticle(dbPool, ctx, res)

	if err != nil {
		log.Printf("Error saving article to db %s: %v\n", res.Article.URL, err)
	}
}
