package response

import "time"

type AuthResponse struct {
	UserId         string
	Token          string
	ExpirationTime time.Time
	HttpStatus     int
}
