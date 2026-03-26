package main

import (
	"log"
	"time"

	"hcmus-news-tele-bot/config"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := config.LoadTelegramBotToken()
	if token == "" {
		log.Fatal("Vui lòng thiết lập TELEGRAM_BOT_TOKEN trong file .env")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Lỗi khởi tạo bot: ", err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		log.Println("Nhận lệnh /start")
		return c.Send("Kết nối bot thành công!")
	})

	log.Println("Bot đang chạy...")
	b.Start()
}
