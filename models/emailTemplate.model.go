package models

type EmailTemplate struct {
	ID      string `gorm:"type:bigint;primary_key" sql:"id"`
	Code    string `gorm:"not null" sql:"code"`
	Subject string `gorm:"not null" sql:"subject"`
	Body    string `gorm:"not null" sql:"body"`
}
