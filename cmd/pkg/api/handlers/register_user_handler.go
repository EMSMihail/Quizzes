package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

//Creating new user via /registration page and transfering data which input in data field into database
//Second fucntion implementing hashing of input password to be stored in database

func RegisterUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if nickname == "" || email == "" || password == "" {
		http.Error(w, "Not all required fields are filled", http.StatusBadRequest)
		return
	}

	user := database.User{
		Nickname:     nickname,
		Email:        email,
		PasswordHash: hashPassword(password),
	}

	err := database.SaveUserToDB(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/success", http.StatusFound)
}

func hashPassword(password string) string {
	// Генерация соли и хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}
