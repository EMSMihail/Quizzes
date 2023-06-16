package handlers

import (
	"fmt"
	"net/http"
)

func SuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Registration Successful!")
}
