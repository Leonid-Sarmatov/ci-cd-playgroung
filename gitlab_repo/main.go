package main

import (
	"fmt"
	"net/http"
)

func hello_kuber_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Bubilda Server!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello_kuber_handler)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
