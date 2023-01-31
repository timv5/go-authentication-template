package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-authentication-template/dto/response"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "No Authorization Header Provided"})
			ctx.Abort()
			return
		}

		claims, err := verifyToken(clientToken)
		if err != nil {
			ctx.JSON(claims.HttpStatus, gin.H{"status": "error", "message": err})
			ctx.Abort()
			return
		}

		ctx.Set("UserId", claims.UserId)
		ctx.Next()
	}
}

func verifyToken(tokenString string) (response.AuthResponse, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return response.AuthResponse{HttpStatus: http.StatusUnauthorized}, err
		} else {
			return response.AuthResponse{HttpStatus: http.StatusBadRequest}, err
		}
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok {
		return response.AuthResponse{HttpStatus: http.StatusUnauthorized}, nil
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return response.AuthResponse{HttpStatus: http.StatusUnauthorized}, nil
	} else {
		return response.AuthResponse{HttpStatus: http.StatusOK, UserId: claims.UserID}, nil
	}
}
