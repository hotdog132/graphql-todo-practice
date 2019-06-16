package models

import "time"

type Event struct {
	ID        string    `json:"id"`
	UserID    string    `json:"text"`
	Done      bool      `json:"done"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
