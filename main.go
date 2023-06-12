package main

import (
	//"fmt"
	//"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "quiz.html")
// 	http.ListenAndServe(":8080", nil)
// }

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "quiz.html")
}

func main() {
	//fmt.Println("Hello, Michael!")
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":5000", nil)
}
