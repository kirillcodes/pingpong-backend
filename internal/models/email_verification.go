package models

import "time"

type EmailVerification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"uniqueIndex;not null"` // один актуальный код на пользователя
	Code      string    `gorm:"size:6;not null"`
	ExpiresAt time.Time `gorm:"index"`
	CreatedAt time.Time
}
