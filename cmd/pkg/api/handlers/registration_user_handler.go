package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if nickname == "" || email == "" || password == "" {
		http.Error(w, "Not all required fields are filled", http.StatusBadRequest)
		return
	}

	existingUser, err := database.GetUserByNicknameOrEmail(db, nickname, email)
	if err != nil {
		log.Println("Error retrieving user:", err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if existingUser != nil {
		http.Error(w, "User with the same nickname or email already exists", http.StatusBadRequest)
		return
	}

	user := database.User{
		Nickname:     nickname,
		Email:        email,
		PasswordHash: hashPassword(password),
	}

	err = database.SaveUserToDB(db, user)
	if err != nil {
		http.Error(w, "Failed to save user to the database", http.StatusInternalServerError)
		log.Println("Error saving user:", err)
		return
	}

	http.Redirect(w, r, "/success", http.StatusFound)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}
