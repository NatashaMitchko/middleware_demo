package main

import (
	"fmt"
	"net/http"
)

func Rendevouz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "The dropoff point is: 123 Wallaby Way, Sydney\n")
}

func main() {
	http.HandleFunc("/", Rendevouz)

	http.ListenAndServe(":8080", nil)
}
