package app

import (
	"database/sql"
	"fmt"
	"malikfajr/todolist-api/helper"

	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {
	DB_HOST := helper.GetEnv("DB_HOST", "127.0.0.1")
	DB_PORT := helper.GetEnv("DB_PORT", "5432")
	DB_USER := helper.GetEnv("DB_USER", "postgres")
	DB_PASSWORD := helper.GetEnv("DB_PASSWORD", "secret")
	DB_NAME := helper.GetEnv("DB_NAME", "belajar")
	SSL_MODE := helper.GetEnv("SSL_MODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, SSL_MODE)

	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	return db
}
