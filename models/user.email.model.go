package models

import "time"

type UserEmail struct {
	ID       string    `gorm:"type:bigint;primary_key" sql:"id"`
	Subject  string    `gorm:"not null" sql:"subject"`
	Body     string    `gorm:"not null" sql:"body"`
	UserID   string    `gorm:"not null" sql:"user_id"`
	DateSent time.Time `gorm:"not null" sql:"date_sent"`
}
