package config

import (
	"hcmus-news-tele-bot/internal/command"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

// LoadTelegramBotToken get telegram bot token from env
func LoadTelegramBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}

// setup bot
func SetupBot(b *tele.Bot, dbPool *pgxpool.Pool) {
	restrictToPrivate := func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat().Type != tele.ChatPrivate {
				return c.Send("Bot này chỉ hỗ trợ chat cá nhân. Vui lòng nhắn tin riêng để sử dụng.")
			}
			return next(c)
		}
	}

	b.Use(restrictToPrivate)

	// register command handlers
	b.Handle("/start", command.Start)
	b.Handle("/subscribe", command.Subscribe(dbPool))
	b.Handle("/unsubscribe", command.Unsubscribe(dbPool))
}
