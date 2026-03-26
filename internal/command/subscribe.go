package command

import (
	"context"
	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

func Subscribe(dbPool *pgxpool.Pool) tele.HandlerFunc {
	return func(c tele.Context) error {
		var new_user = model.User{
			ID:           strconv.FormatInt(c.Sender().ID, 10),
			Name:         c.Sender().FirstName,
			IsSubscribed: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := repository.SaveUser(dbPool, ctx, new_user)

		if err != nil {
			log.Printf("Lỗi đăng ký user %s: %v", new_user.ID, err)
			return c.Send("Đã có lỗi xảy ra khi đăng ký. Vui lòng thử lại sau.")
		}

		msg := "Bạn đã đăng ký thành công!, id: " + new_user.ID
		msg += "\nBạn sẽ nhận được thông báo khi có tin tức mới từ các trang web của HCMUS."
		return c.Send(msg)
	}
}
