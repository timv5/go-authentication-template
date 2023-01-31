package service

import (
	"github.com/golang-jwt/jwt/v4"
	"go-authentication-template/configs"
	"go-authentication-template/dto/response"
	"go-authentication-template/models"
	"net/http"
	"time"
)

type Authentication interface {
	GenerateToken(user *models.User) (response.AuthResponse, error)
	VerifyToken(tokenString string) (response.AuthResponse, error)
}

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

type AuthService struct {
	conf *configs.Config
}

func NewAuthService(conf *configs.Config) *AuthService {
	return &AuthService{conf: conf}
}

func (auth *AuthService) VerifyToken(tokenString string) (response.AuthResponse, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.conf.JwtSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return response.AuthResponse{HttpStatus: http.StatusUnauthorized}, err
		} else {
			return response.AuthResponse{HttpStatus: http.StatusBadRequest}, err
		}
	}

	if !tkn.Valid {
		return response.AuthResponse{HttpStatus: http.StatusUnauthorized}, nil
	} else {
		return response.AuthResponse{HttpStatus: http.StatusOK}, nil
	}
}

func (auth *AuthService) GenerateToken(user *models.User) (response.AuthResponse, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(auth.conf.JwtExpiration))
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	secret := []byte(auth.conf.JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(secret)
	if err != nil {
		return response.AuthResponse{HttpStatus: http.StatusBadRequest}, err
	} else {
		return response.AuthResponse{UserId: user.ID, Token: stringToken, ExpirationTime: expirationTime, HttpStatus: http.StatusOK}, nil
	}
}
