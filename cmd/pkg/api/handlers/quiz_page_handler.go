package handlers

import (
	"fmt"
	"net/http"
)

func QuizPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
