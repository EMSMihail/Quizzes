package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
)

func RegistrationPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

	// tmpl, err := template.ParseFiles("../../web/templates/registration.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	tmplDir := os.Getenv("TEMPLATES_DIR")
	if tmplDir == "" {
		log.Fatal("TEMPLATES_DIR environment variable is not set")
	}

	tmplPath := filepath.Join(tmplDir, "registration.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
