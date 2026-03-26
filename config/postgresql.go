package config

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// config
var (
	pool *pgxpool.Pool // PostgreSQL connection pool
	once sync.Once     // ensure singleton instance
)

func ConnectPostgreSQL() {
	once.Do(func() {
		dbUrl := os.Getenv("SUPABASE_URL") // get dbUrl from env file

		if dbUrl == "" {
			log.Fatal("SUPABASE_URL is not set in .env file")
		}

		config, err := pgxpool.ParseConfig(dbUrl)
		if err != nil {
			log.Fatalf("Unable to parse SUPABASE_URL: %v", err)
		}

		config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

		pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		// verify connection
		if err := pool.Ping(context.Background()); err != nil {
			log.Fatalf("Unable to connect to database: %v", err)
		}

		log.Println("Connected to PostgreSQL database successfully")

	})
}

// get pool
func GetPool() *pgxpool.Pool {
	if pool == nil {
		log.Fatal("Database connection not initialized. Call ConnectPostgreSQL() first.")
	}
	return pool
}

func CloseDB() {
	if pool != nil {
		pool.Close()
	}
}
