package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type PageData struct {
	Title    string
	Message  string
	IsError  bool
	Nickname string
	Email    string
	Password string
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if nickname == "" || email == "" || password == "" {
		showRegistrationPage(w, "Not all required fields are filled", true, nickname, email, password)
		return
	}

	existingUser, err := database.GetUserByNicknameOrEmail(db, nickname, email)
	if err != nil {
		log.Println("Error retrieving user:", err)
		showRegistrationPage(w, "Error retrieving user", true, nickname, email, password)
		return
	}

	if existingUser != nil {
		showRegistrationPage(w, "User with the same nickname or email already exists", true, nickname, email, password)
		return
	}

	user := database.User{
		Nickname:     nickname,
		Email:        email,
		PasswordHash: hashPassword(password),
	}

	err = database.SaveUserToDB(db, user)
	if err != nil {
		log.Println("Error saving user:", err)
		showRegistrationPage(w, "Failed to save user to the database", true, nickname, email, password)
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

func showRegistrationPage(w http.ResponseWriter, message string, isError bool, nickname string, email string, password string) {
	data := PageData{
		Title:    "Registration Page",
		Message:  message,
		IsError:  isError,
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	tmpl, err := template.ParseFiles("../../web/templates/registration.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
