package repository

import (
	"context"
	"hcmus-news-tele-bot/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveUser(
	db *pgxpool.Pool,
	ctx context.Context,
	user model.User,
) error {
	query := `
		insert into users (id, name, is_subscribed, created_at, updated_at)
		values ($1, $2, $3, $4, $5)
		on conflict (id) do update set
			name = excluded.name,
			is_subscribed = excluded.is_subscribed,
			updated_at = excluded.updated_at
	`

	_, err := db.Exec(ctx, query, user.ID, user.Name, user.IsSubscribed, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
