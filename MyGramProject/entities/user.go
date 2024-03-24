package entities

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
