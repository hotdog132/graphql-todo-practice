package models

import "time"

type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
