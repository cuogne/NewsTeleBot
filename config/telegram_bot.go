package config

import "os"

// LoadTelegramBotToken get telegram bot token from env
func LoadTelegramBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}
