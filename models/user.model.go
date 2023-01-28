package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:bigint;primary_key" sql:"id"`
	Email     string    `gorm:"not null" sql:"email"`
	Password  string    `gorm:"not null" sql:"password"`
	Username  string    `gorm:"not null" sql:"username"`
	CreatedAt time.Time `gorm:"not null" sql:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" sql:"updatedAt"`
}
