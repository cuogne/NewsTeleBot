package repository

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetArticle(
	db *pgxpool.Pool,
	ctx context.Context,
	category string,
) ([]model.Article, error) {

	query := ""
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

	var articlesList []model.Article
	for rows.Next() {
		var n model.Article
		if err := rows.Scan(&n.Title, &n.URL); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		articlesList = append(articlesList, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return articlesList, nil
}
