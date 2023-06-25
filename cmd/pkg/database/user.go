package database

import (
	"database/sql"
)

//Declaration of struct User with 4 fields
//First func shows all user's table entries via SELECT
//Second func creating new entry in user's table via INSERT

type User struct {
	ID           int
	Nickname     string
	Email        string
	PasswordHash string
}

func GetUsersFromDB(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUserToDB(db *sql.DB, user User) error {
	query := `INSERT INTO users (nickname, email, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, user.Nickname, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}
