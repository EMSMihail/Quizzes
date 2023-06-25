package handlers

import (
	"fmt"
	"net/http"
)

//Just output on the /success page text of finalizing of registration

func SuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Successful!")
}

func RegistrationSuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Registration Successful!")
}
