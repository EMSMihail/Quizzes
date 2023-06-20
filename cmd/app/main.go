package main

import (
	"log"
	"net/http"

	"github.com/EMSMihail/Quizzes/cmd/pkg/api/handlers"
	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database!")

	users, err := database.GetUsersFromDB(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Users:")

	for _, user := range users {
		log.Println("ID:", user.ID, "| Username:", user.Nickname, "| E-Mail:", user.Email, "| Password_hash:", user.PasswordHash)
	}

	fs := http.FileServer(http.Dir("../../web/static/css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login", handlers.LoginPageHandler)
	http.HandleFunc("/registration", handlers.RegistrationPageHandler)
	http.HandleFunc("/success", handlers.SuccessPageHandler)
	http.HandleFunc("/", handlers.QuizPageHandler)
	//http.Handle("")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
