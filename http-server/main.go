package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		log.Fatal("PORT env var is missing.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.ListenAndServe(fmt.Sprintf(":%v", p), nil)
}
