package database

import (
	"database/sql"
)

func GetHashPassFromDB(db *sql.DB, email string) (string, error) {
	var hashPass string
	err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashPass)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь с указанным email не найден
			return "", nil
		}
		return "", err
	}
	return hashPass, nil
}
