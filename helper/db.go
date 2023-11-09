package helper

import "database/sql"

func NewDb() *sql.DB {
	connStr := "user=postgres dbname=belajar password='secret' host='127.0.0.1' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	PanicIfError(err)

	return db
}
