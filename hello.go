package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/zc", hello)

	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Docker Form Golang!")

}
