package command

import (
	"context"
	"hcmus-news-tele-bot/internal/repository"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

func Unsubscribe(dbPool *pgxpool.Pool) tele.HandlerFunc {
	return func(c tele.Context) error {
		user_id := strconv.FormatInt(c.Sender().ID, 10)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := repository.UnsubscribeUser(dbPool, ctx, user_id)
		if err != nil {
			log.Printf("Lỗi hủy đăng ký user %s: %v", user_id, err)
			return c.Send("Có lỗi xảy ra khi hủy đăng ký nhận thông báo, vui lòng thử lại sau!")
		}

		msg := "Bạn đã hủy đăng ký nhận thông báo thành công!"
		return c.Send(msg)
	}
}
