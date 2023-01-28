package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"go-authentication-template/configs"
)

func NewConnection(config *configs.Configurations, logger hclog.Logger) (*sqlx.DB, error) {
	var conn string

	if config.DBConn != "" {
		conn = config.DBConn
	} else {
		host := config.DBHost
		username := config.DBUser
		password := config.DBPass
		port := config.DBPort
		name := config.DBName
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, name, password)
	}

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	} else {
		return db, err
	}
}
