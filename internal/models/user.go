package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;not null"` // уникальный индекс на email
	CreatedAt time.Time
	UpdatedAt time.Time
}
