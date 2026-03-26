package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UnsubscribeUser(
	db *pgxpool.Pool,
	ctx context.Context,
	userID string,
) error {
	query := `
		update users
		set is_subscribed = false, updated_at = current_timestamp
		where id = $1
	`
	_, err := db.Exec(ctx, query, userID)
	return err
}
