# go-authentication-template
Authentication service written in Golang connecting to Postgres database. Includes login, register and 2 example protected routes.
After a successful registration email is sent to a user.

## Used technology and prerequisites
- Golang
- Docker, Docker Compose
- Postgres

## How to run the project?
- go to /docker_db
- run command: docker-compose -f docker-compose.yml up -d
- generate initial schema: generate_schema.sql
- insert initial data: init_migration.sql

## Endpoints and functionality
- root url: http://localhost:8080

### Register
- unprotected route
- (POST) http://localhost:8080/api/auth/register
- request:
  {
  "email": "test@gmail.com",
  "password": "password",
  "passwordConfirm": "password",
  "username": "username_t"
  }
- response:
{
    "user": {
    "ID": 0,
    "Email": "test@gmail.com",
    "Username": "username_t"
    }
    }

### Login
- unprotected route
- (POST) http://localhost:8080/api/auth/login
- request:
  {
  "email": "test@gmail.com",
  "password": "password"
  }
- response:
  {
  "payload": {
  "UserId": "e12a74ca-006b-4dd7-807b-e93ec81b6618",
  "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJlMTJhNzRjYS0wMDZiLTRkZDctODA3Yi1lOTNlYzgxYjY2MTgiLCJleHAiOjE2NzUxMTYwNDB9.PeqlPwksHopSVZUm2ckAihZnKARTQzvfU6iH_7rhZF4",
  "ExpirationTime": "2023-01-30T23:00:40.924353547+01:00"
  },
  "status": "success"
  }


### getByEmail
- protected route
- (POST) http://localhost:8080/api/users/getByEmail
- request:
  {
  "email": "test@test.com"
  }
- response: 
 {
  "ID": 1000,
  "Email": "test@test.com",
  "Password": "test_password",
  "Username": "test_username",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z"
  }

### getByUsername
- protected route
- (GET) http://localhost:8080/api/users/getByUsername?username=test_username
- response:
  {
  "ID": 1000,
  "Email": "test@test.com",
  "Password": "test_password",
  "Username": "test_username",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z"
  }
