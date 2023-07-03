package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/EMSMihail/Quizzes/cmd/pkg/database"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Title": "Login Page",
		"Error": r.URL.Query().Get("error"),
	}

	if r.URL.Query().Get("error") == "1" {
		data["Error"] = "Invalid username or password"
	}

	//tmpl, err := template.ParseFiles("../../web/templates/login.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	tmplDir := os.Getenv("TEMPLATES_DIR")
	if tmplDir == "" {
		log.Fatal("TEMPLATES_DIR environment variable is not set")
	}

	tmplPath := filepath.Join(tmplDir, "login.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		db, err := database.ConnectToDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		LoginUserHandler(w, r, db)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
