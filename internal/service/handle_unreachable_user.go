package service

import (
	"context"
	"errors"
	"hcmus-news-tele-bot/internal/repository"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

// isUserUnreachable checks if the user blocked the bot or deactivated their account,
// meaning further sending attempts are pointless.
func isUserUnreachable(err error) bool {
	return errors.Is(err, tele.ErrBlockedByUser) ||
		errors.Is(err, tele.ErrUserIsDeactivated) ||
		errors.Is(err, tele.ErrNotStartedByUser)
}

func autoUnsubscribe(dbPool *pgxpool.Pool, userID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repository.UnsubscribeUser(dbPool, ctx, userID); err != nil {
		log.Printf("Error auto-unsubscribing user %s: %v\n", userID, err)
		return
	}

	log.Printf("User %s blocked/deactivated the bot, auto-unsubscribed", userID)
}
