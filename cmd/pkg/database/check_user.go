package database

import (
	"database/sql"
)

//Declaration of struct User with 4 fields
//First func shows all user's table entries via SELECT
//Second func creating new entry in user's table via INSERT

//	func GetHashPassFromDB(db *sql.DB, email string) {
//		var hash string
//		err := db.QueryRow("SELECT password_hash FROM users WHERE email=$1", email).Scan(&hash)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
func GetHashPassFromDB(db *sql.DB, email string) (string, error) {
	var hashedPass string
	err := db.QueryRow("SELECT password_hash FROM users WHERE email = $1", email).Scan(&hashedPass)
	if err != nil {
		return "", err
	}
	return hashedPass, nil
}
