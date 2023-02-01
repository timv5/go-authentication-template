package repository

import (
	"go-authentication-template/models"
	"gorm.io/gorm"
)

type EmailTemplateRepositoryInterface interface {
	GetByCode(code string) models.EmailTemplate
}

type EmailTemplateRepository struct {
	postgresDB *gorm.DB
}

func NewEmailTemplateRepository(postgresDB *gorm.DB) *EmailTemplateRepository {
	return &EmailTemplateRepository{postgresDB: postgresDB}
}

func (repo EmailTemplateRepository) GetByCode(code string) models.EmailTemplate {
	var template models.EmailTemplate
	repo.postgresDB.Raw("select id, code, subject, body from email_template where code = ?", code).Scan(&template)
	if (models.EmailTemplate{} == template) {
		return models.EmailTemplate{}
	} else {
		return template
	}
}
