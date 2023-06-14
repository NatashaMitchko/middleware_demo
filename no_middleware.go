package main

import (
	"errors"
	"fmt"
	"net/http"
)

func BadSecurity(password string) error {
	if password != "007" {
		return errors.New("wrong passcode! Our system has been compromised!")
	}
	return nil
}

func Rendevouz(w http.ResponseWriter, req *http.Request) {
	// pw := req.Header.Get("Password")
	// TODO: Check the password here!
	// if err != nil return 403 unauthorized

	fmt.Fprintf(w, "The dropoff point is: 42 Wallaby Way, Sydney\n")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/rendezvous", Rendevouz)

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
