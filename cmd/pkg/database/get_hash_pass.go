package database

import (
	"database/sql"
)

func GetHashPassFromDB(db *sql.DB, email string) (string, error) {
	var hashedPass string
	err := db.QueryRow("SELECT password_hash FROM users WHERE email = $1", email).Scan(&hashedPass)
	if err != nil {
		return "", err
	}
	return hashedPass, nil
}
