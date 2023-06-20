package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassFromDB, err := database.GetHashPassFromDB(db, email)
	if err != nil {
		log.Fatal(err)
	}

	if hashedPassFromDB == "" {
		fmt.Fprintf(w, "Error: Invalid email or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassFromDB), []byte(password))
	if err != nil {
		fmt.Fprintf(w, "Error: Invalid email or password")
		return
	}

	fmt.Fprintf(w, "Success: Passwords match")

	// // Здесь вы можете добавить свою логику проверки учетных данных пользователя
	// if email == "test@test.test" && password == "1234" {
	// 	// Успешный вход
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// } else {
	// 	// Неверные учетные данные
	// 	http.Redirect(w, r, "/login?error=1", http.StatusFound)
	// }
}
