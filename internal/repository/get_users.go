package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetSubscribedUsers(db *pgxpool.Pool, ctx context.Context) ([]string, error) {
	query := `select id from users where is_subscribed = true`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		users = append(users, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
