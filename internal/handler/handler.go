package handler

import (
	"context"
	"fmt"
	"log"

	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"hcmus-news-tele-bot/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robfig/cron/v3"
	tele "gopkg.in/telebot.v4"
)

func StartCronJob(b *tele.Bot, dbPool *pgxpool.Pool) {
	c := cron.New()
	_, err := c.AddFunc("@every 30s", func() {
		log.Println("starting cron cycle...")
		runCronCycle(b, dbPool)
	})
	if err != nil {
		log.Fatalf("Lỗi thiết lập cron job: %v", err)
	}

	log.Println("cron job is working")
	c.Start()
}

func runCronCycle(b *tele.Bot, dbPool *pgxpool.Pool) {
	// crawl all news
	articles, err := service.GetArticles()
	if err != nil {
		log.Printf("Lỗi crawl cron: %v\n", err)
		return
	}

	// filter new articles
	newArticles := service.FilterNewArticles(dbPool, articles)
	if len(newArticles) == 0 {
		log.Println("Không có bài viết mới.")
		return
	}

	fmt.Println(newArticles)

	jobs := make(chan model.SummaryJob, len(newArticles))
	results := make(chan model.SummaryResult, len(newArticles))

	// create 3 worker pool
	for w := 1; w <= 3; w++ {
		go service.Worker(w, jobs, results)
	}

	for _, a := range newArticles {
		jobs <- a
	}
	close(jobs)

	// get subscribed users
	users, err := repository.GetSubscribedUsers(dbPool, context.Background())
	if err != nil {
		log.Printf("Error while getting subscribed users: %v\n", err)
	}

	for range newArticles {
		res := <-results                // receive summary from gemini
		service.SendTele(b, users, res) // send for subscribed users
		service.SaveToDB(dbPool, res)   // save to db
	}
}
