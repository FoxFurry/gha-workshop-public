package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request at /hello")
	fmt.Fprintf(w, "Hello workshop\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Starting the server at :8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}
