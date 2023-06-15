package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// func mainPageHandler(w http.ResponseWriter, r *http.Request) {
// 	//http.ServeFile(w, r, "/web/templates/login.html")
// 	tmpl, err := http.Dir("../../web/templates").Open("login.html")
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer tmpl.Close()

// 	// Serve the HTML template
// 	fi, err := tmpl.Stat()
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	http.ServeContent(w, r, "login.html", fi.ModTime(), tmpl)
// }

// func cssHandler(w http.ResponseWriter, r *http.Request) {
// 	filePath := "../../web/static/login.css"
// 	defer func() {
// 		if err := recover(); err != nil {
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			// You can also log the error or take any other appropriate action
// 		}
// 	}()

// 	http.ServeFile(w, r, filePath)
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]interface{}{
// 		"Title": "Login Page",
// 	}

// 	tmpl, err := template.ParseFiles("../../web/templates/login.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := tmpl.Execute(w, data); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func registrationHandler(w http.ResponseWriter, r *http.Request) {
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

func connectToDB() (*sql.DB, error) {
	connStr := "postgres://aizek:1234@localhost/quizzes?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключение к базе данных успешно!")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nickname, email, password_hash string
		err := rows.Scan(&id, &nickname, &email, &password_hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "| Username:", nickname, "| E-Mail:", email, "| Password_hash:", password_hash)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("../../web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.HandleFunc("/", loginHandler)
	http.HandleFunc("/", registrationHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

// func main() {
// 	//fmt.Println("Hello, Michael!")
// 	http.HandleFunc("/", mainPageHandler)
// 	http.HandleFunc("/static/css/login.css", cssHandler)

// 	fs := http.FileServer(http.Dir("../../web/static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	//http.ListenAndServe(":5000", nil)
// 	//http.HandleFunc("/", mainPageHandler)
// 	// http.HandleFunc("../../web/static/login.css", cssHandler)
// 	http.ListenAndServe(":5000", nil)
// }

// func main() {
// 	http.HandleFunc("/", mainPageHandler)

// 	fs := http.FileServer(http.Dir("../../web/static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	http.HandleFunc("/static/css/login.css", cssHandler)

// 	http.ListenAndServe(":5000", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":5000", nil)
// }
