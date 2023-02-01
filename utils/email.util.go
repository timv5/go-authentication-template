package utils

import (
	"crypto/tls"
	"go-authentication-template/configs"
	"go-authentication-template/models"
	"go-authentication-template/repository"
	"gopkg.in/gomail.v2"
	"log"
)

type Email struct {
	Subject string
	Body    string
}

func SendEmail(user *models.User, emailTemplateRepo *repository.EmailTemplateRepository,
	userEmailRepository *repository.UserEmailTemplateRepository, configs *configs.Config) {
	var emailTemplate models.EmailTemplate
	emailTemplate = emailTemplateRepo.GetByCode("register")
	if (models.EmailTemplate{} == emailTemplate) {
		log.Println("Email template not found")
		return
	}

	email, err := userEmailRepository.Save(emailTemplate.Subject, emailTemplate.Body, user.ID)
	if err != nil {
		log.Println("Email cannot be sent")
		return
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "ourplatform@test.com")
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", email.Subject)
	message.SetBody("text/html", email.Body)

	d := gomail.NewDialer(configs.DBHost, configs.SMTPPort, configs.SMTPUser, configs.SMTPPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(message); err != nil {
		log.Println("Email cannot be sent: ", err)
	}

}
