package handlers

import (
	"database/sql"
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
		// Пользователь не существует в базе данных
		http.Redirect(w, r, "/login?error=1", http.StatusFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassFromDB), []byte(password))
	if err != nil {
		// Неверный email или пароль
		http.Redirect(w, r, "/login?error=1", http.StatusFound)
		return
	}

	// Успешная аутентификация
	http.Redirect(w, r, "/success", http.StatusFound)
}
