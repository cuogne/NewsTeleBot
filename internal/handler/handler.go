package handler

import (
	"context"
	"log"
	"sync/atomic"

	"hcmus-news-tele-bot/internal/model"
	"hcmus-news-tele-bot/internal/repository"
	"hcmus-news-tele-bot/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robfig/cron/v3"
	tele "gopkg.in/telebot.v4"
)

var cronCycleRunning atomic.Bool

func StartCronJob(b *tele.Bot, dbPool *pgxpool.Pool) {
	c := cron.New()
	schedule := "@every 10m" // every 10 minutes

	_, err := c.AddFunc(schedule, func() {
		runCronCycle(b, dbPool)
	})
	if err != nil {
		log.Fatalf("Error starting cron job: %v", err)
	}

	log.Println("cron job is working, 10 minutes per cycle")
	c.Start()
}

func runCronCycle(b *tele.Bot, dbPool *pgxpool.Pool) {
	if !cronCycleRunning.CompareAndSwap(false, true) {
		log.Println("Skip cron tick: previous cycle is still running")
		return
	}
	defer cronCycleRunning.Store(false)

	// get subscribed users first
	users, err := repository.GetSubscribedUsers(dbPool, context.Background())
	if err != nil {
		// no user -> no need to crawl and send
		log.Printf("Error while getting subscribed users: %v\n", err)
		return
	}

	// check user count
	if len(users) == 0 {
		log.Println("No subscribed users found")
		return
	}

	// crawl all news
	articles, err := service.GetArticles()
	if err != nil {
		log.Printf("Error crawling cron: %v\n", err)
		return
	}

	// filter new articles
	newArticles, err := service.FilterNewArticles(dbPool, articles)
	if err != nil {
		log.Printf("Error filtering articles: %v\n", err)
	}

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
		log.Println("[NEWS]: ", a.Article.Title)
		jobs <- a
	}
	close(jobs)

	for range newArticles {
		res := <-results                // receive summary from gemini
		service.SendTele(b, users, res) // send for subscribed users
		service.SaveToDB(dbPool, res)   // save to db
	}
}
