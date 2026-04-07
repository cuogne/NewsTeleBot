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

	table := category // define table name = category
	query := fmt.Sprintf("select title, url from %s", table)

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
