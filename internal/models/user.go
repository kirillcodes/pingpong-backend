package models

import "time"

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Email          string `gorm:"uniqueIndex;not null"` // уникальный индекс на email
	Username       string `gorm:"size:100;not null"`
	PasswordHash   string `gorm:"not null"`
	EmailConfirmed bool   `gorm:"default:false"`
	ConfirmCode    string `gorm:"size:6"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
