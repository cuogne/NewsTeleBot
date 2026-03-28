package handler

import (
	"context"
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
	time := "@every 10m" // every 10 minutes

	_, err := c.AddFunc(time, func() {
		runCronCycle(b, dbPool)
	})
	if err != nil {
		log.Fatalf("Error starting cron job: %v", err)
	}

	log.Println("cron job is working, 10 minutes per cycle")
	c.Start()
}

func runCronCycle(b *tele.Bot, dbPool *pgxpool.Pool) {
	// crawl all news
	articles, err := service.GetArticles()
	if err != nil {
		log.Printf("Error crawling cron: %v\n", err)
		return
	}

	// filter new articles
	newArticles := service.FilterNewArticles(dbPool, articles)
	if len(newArticles) == 0 {
		return
	}

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
