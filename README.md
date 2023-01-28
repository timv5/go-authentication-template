# go-authentication-template
Authentication service written in Golang

## source
- https://medium.com/swlh/building-a-user-auth-system-with-jwt-using-golang-30892659cc0
- https://github.com/d-vignesh/go-jwt-auth

# Prerequisites
- go get github.com/julienschmidt/httprouter
- go get github.com/jinzhu/gorm
- go get github.com/dgrijalva/jwt-go
- go get github.com/gorilla/mux


## getByEmail
- response: 
 {
  "ID": 1000,
  "Email": "test@test.com",
  "Password": "test_password",
  "Username": "test_username",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z"
  }

## getByUsername
- response:
  {
  "ID": 1000,
  "Email": "test@test.com",
  "Password": "test_password",
  "Username": "test_username",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z"
  }

## register
- response:
  {
  "user": {
  "ID": 0,
  "Email": "test@gmail.com",
  "Username": "username_t"
  }
  }