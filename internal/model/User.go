package model

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	IsSubscribed bool      `json:"is_subscribed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
