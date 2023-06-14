package main

import (
	"fmt"
	"net/http"
)

//func BadSecurity TODO: complete the middleware function

func RendevouzMW(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "The dropoff point is: 42 Wallaby Way, Sydney\n")
}

func Enemies(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Voldemort, Darth Vader, Hans Gruber\n")
}

func SecureLine(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You can call the president at: (646) 504-8370\n")
}
func main() {

	mux := http.NewServeMux()

	// TODO: add middleware

	mux.HandleFunc("/rendezvous", RendevouzMW)
	mux.HandleFunc("/enemy-spies", Enemies)
	mux.HandleFunc("/secure-line", SecureLine)

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
