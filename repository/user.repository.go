package repository

import (
	uuid "github.com/satori/go.uuid"
	"go-authentication-template/models"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryInterface interface {
	SaveUser(email string, password string, username string) (models.User, error)
	GetUserByEmail(email string) models.User
	GetUserByUsername(username string) models.User
}

type UserRepository struct {
	postgresDB *gorm.DB
}

func NewUserRepository(postgresDB *gorm.DB) *UserRepository {
	return &UserRepository{postgresDB: postgresDB}
}

func (repo *UserRepository) SaveUser(email string, password string, username string) (models.User, error) {
	nowTime := time.Now()
	createUser := models.User{
		ID:        uuid.NewV4().String(),
		Email:     email,
		Password:  password,
		Username:  username,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}

	savedUser := repo.postgresDB.Create(&createUser)
	if savedUser.Error != nil {
		return models.User{}, savedUser.Error
	} else {
		return createUser, nil
	}
}

func (repo *UserRepository) GetUserByEmail(email string) models.User {
	var user models.User
	repo.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where email = ?",
		email).Scan(&user)
	if (models.User{} == user) {
		return models.User{}
	} else {
		return user
	}
}

func (repo *UserRepository) GetUserByUsername(username string) models.User {
	var user models.User
	repo.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where username = ?",
		username).Scan(&user)
	if (models.User{} == user) {
		return models.User{}
	} else {
		return user
	}
}
