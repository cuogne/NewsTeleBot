package repository

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetNews(
	db *pgxpool.Pool,
	ctx context.Context,
	category string,
) ([]model.News, error) {
	var query string
	switch category {
	case "fithcmus":
		query = `select title, url from fitnews`
	case "lichthi":
		query = `select title, url from lichthi`
	case "thongbao":
		query = `select title, url from thongbaopkt`
	case "hcmus":
		query = `select title, url from hcmus`
	default:
		return nil, fmt.Errorf("unsupported category: %s", category)
	}

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var newsList []model.News
	for rows.Next() {
		var n model.News
		if err := rows.Scan(&n.Title, &n.URL); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		newsList = append(newsList, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return newsList, nil
}
