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

func GetUserByNicknameOrEmail(db *sql.DB, nickname string, email string) (*User, error) {
	query := `SELECT nickname, email, password_hash FROM users WHERE nickname = $1 OR email = $2 LIMIT 1`

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
