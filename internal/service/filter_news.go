package service

import (
	"context"
	"errors"
	"fmt"
	"hcmus-news-tele-bot/config"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FilterNewArticles(
	dbPool *pgxpool.Pool,
	articles []model.Article,
) ([]model.SummaryJob, error) {
	// categories := []string{"fithcmus", "lichthi", "thongbao", "hcmus"} // luoi qa nen set cung data
	categories := make([]string, len(config.Feeds))
	for i, feed := range config.Feeds {
		categories[i] = feed.Category
	}

	existUrls := make(map[string]bool)
	var lookupErrors []error

	for _, category := range categories {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		dataArticles, err := repository.GetArticle(dbPool, ctx, category)
		cancel() // ensure context is canceled after use

		/*
			catch error and skip category,
			do not return immediately to avoid losing all new articles
			if one category has db issue
		*/
		if err != nil {
			log.Printf("Error in db %s: %v\n", category, err)
			lookupErrors = append(lookupErrors, fmt.Errorf("category %s: %w", category, err))
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

	if len(lookupErrors) > 0 {
		return nil, errors.Join(lookupErrors...)
	}

	return newArticles, nil
}
