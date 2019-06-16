package models

import "time"

// User user
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
