package main

import (
	"fmt"
	"net/http"
)

func BadSecurity(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Password") != "007" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Wrong Passcode\n"))
			return
		}
		handler.ServeHTTP(w, req)
	})
}

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
	wrappedmux := BadSecurity(mux)

	mux.HandleFunc("/rendezvous", RendevouzMW)
	mux.HandleFunc("/enemy-spies", Enemies)
	mux.HandleFunc("/secure-line", SecureLine)

	s := http.Server{
		Addr:    ":8080",
		Handler: wrappedmux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
