package service

import (
	"context"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveToDB(dbPool *pgxpool.Pool, res model.SummaryResult) {
	if res.Summary == "" {
		return
	}

	// save to db
	err := repository.SaveArticle(
		dbPool,
		context.Background(),
		res,
	)

	if err != nil {
		log.Printf("Error saving article to db %s: %v\n", res.Article.URL, err)
	}
}
