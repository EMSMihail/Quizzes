package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Function for connecting to postgresql database with user=aizek, pass=1234, db=quizzes

func ConnectToDB() (*sql.DB, error) {
	//host := "localhost"
	// host := "host.docker.internal"
	host := "quizzes_db"
	//port := 5433 // Порт, привязанный к контейнеру (5433 в данном случае)
	port := 5432
	user := "aizek"
	password := "1234"
	dbname := "quizzes"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func ConnectToDB() (*sql.DB, error) {
// 	connStr := "postgres://aizek:1234@localhost/quizzes?sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func GetUserByNicknameOrEmail(db *sql.DB, nickname string, email string) (*User, error) {
	query := `SELECT nickname, email, password FROM users WHERE nickname = $1 OR email = $2 LIMIT 1`

	row := db.QueryRow(query, nickname, email)

	user := User{}

	err := row.Scan(&user.Nickname, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
