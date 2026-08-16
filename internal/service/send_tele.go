package service

import (
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"html"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

func SendTele(b *tele.Bot, dbPool *pgxpool.Pool, users []string, res model.SummaryResult) {
	title := html.EscapeString(strings.TrimSpace(res.Article.Title))
	summary := html.EscapeString(strings.TrimSpace(res.Summary))
	link := html.EscapeString(res.Article.URL)

	msg := ""

	if summary == "" {
		// summary empty -> "lichthi" and "thongbao" (from worker)
		msg = fmt.Sprintf("📰 <b>[TIN MỚI]: %s</b>\n\nChi tiết xem tại: %s", title, link)
	} else {
		msg = fmt.Sprintf("📰 <b>[TIN MỚI]: %s</b>\n\n%s\n\nChi tiết xem tại: %s", title, summary, link)
	}

	for _, uID := range users {
		chatID, err := strconv.ParseInt(uID, 10, 64)
		if err != nil {
			continue
		}
		recipient := &tele.Chat{ID: chatID}
		_, err = b.Send(recipient, msg, tele.ModeHTML)
		if err != nil {
			log.Printf("Error sending message to user %s: %v\n", uID, err)

			if isUserUnreachable(err) {
				autoUnsubscribe(dbPool, uID)
			}
		}
	}
}
