package repository

import (
	uuid "github.com/satori/go.uuid"
	"go-authentication-template/models"
	"gorm.io/gorm"
	"time"
)

type UserEmailRepositoryInterface interface {
	Save(subject string, body string, userId string) (models.UserEmail, error)
}

type UserEmailTemplateRepository struct {
	postgresDB *gorm.DB
}

func NewUserEmailTemplateRepository(postgresDB *gorm.DB) *UserEmailTemplateRepository {
	return &UserEmailTemplateRepository{postgresDB: postgresDB}
}

func (repo UserEmailTemplateRepository) Save(subject string, body string, userId string) (models.UserEmail, error) {
	nowTime := time.Now()

	createEmail := models.UserEmail{
		ID:       uuid.NewV4().String(),
		Subject:  subject,
		Body:     body,
		UserID:   userId,
		DateSent: nowTime,
	}

	savedEmail := repo.postgresDB.Create(&createEmail)
	if savedEmail.Error != nil {
		return models.UserEmail{}, savedEmail.Error
	} else {
		return createEmail, nil
	}
}
