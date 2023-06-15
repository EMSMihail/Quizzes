package main

import (
	//"fmt"
	//"database/sql"
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

func main() {
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
