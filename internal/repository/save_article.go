package repository

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveArticle(
	db *pgxpool.Pool,
	ctx context.Context,
	res model.SummaryResult,
) error {

	table := res.Category   // variable to hold table query
	numArticlesToKeep := 20 // store 20 articles per category

	query := fmt.Sprintf(`
		insert into %s (url, title, send_at, prompt_token, completion_token)
		values ($1, $2, current_timestamp, $3, $4)
		on conflict (url) do nothing
	`, table)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction for %s: %w", table, err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query,
		res.Article.URL,
		res.Article.Title,
		res.PromptToken,
		res.CompletionToken,
	)

	if err != nil {
		return fmt.Errorf("failed to insert article to %s: %w", table, err)
	}

	pruneQuery := fmt.Sprintf(`
		delete from %s
		where url in (
			select url from %s
			order by send_at desc nulls last, url desc
			offset $1
		)
	`, table, table)

	_, err = tx.Exec(ctx, pruneQuery, numArticlesToKeep)
	if err != nil {
		return fmt.Errorf("failed to prune old articles in %s: %w", table, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction for %s: %w", table, err)
	}

	return nil
}
