package service

import (
	"context"
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FilterNewArticles(
	dbPool *pgxpool.Pool,
	articles []model.Article,
) []model.SummaryJob {
	// categories := []string{"fithcmus", "lichthi", "thongbao", "hcmus"} // luoi qa nen set cung data
	categories := make([]string, len(config.Feeds))
	for i, feed := range config.Feeds {
		categories[i] = feed.Category
	}

	existUrls := make(map[string]bool)

	for _, category := range categories {
		dataArticles, err := repository.GetArticle(dbPool, context.Background(), category)
		if err != nil {
			log.Printf("Error in db %s: %v\n", category, err)
			continue
		}
		for _, dn := range dataArticles {
			existUrls[dn.URL] = true
		}
	}

	var newArticles []model.SummaryJob
	for _, a := range articles {
		if !existUrls[a.URL] {
			newArticles = append(newArticles, model.SummaryJob{
				Article:  a,
				Category: a.Category,
			})
		}
	}

	return newArticles
}
