package handlers

import (
	"html/template"
	"net/http"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
)

//Function that implements functionality of /registration page and some checks which needed for this page to work properly

func RegistrationPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// If the request method is POST, call the registration handler
		db, err := database.ConnectToDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		RegisterUserHandler(w, r, db)
		return
	}

	data := map[string]interface{}{
		"Title": "Registration Page",
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
