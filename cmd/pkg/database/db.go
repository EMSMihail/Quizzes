package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Function for connecting to postgresql database with user=aizek, pass=1234, db=quizzes

func ConnectToDB() (*sql.DB, error) {
	connStr := "postgres://aizek:1234@localhost/quizzes?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
