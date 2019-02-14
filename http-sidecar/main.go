package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	fp := os.Getenv("FE_PORT")
	if fp == "" {
		log.Fatal("FE_PORT env var is missing.")
	}

	bp := os.Getenv("BE_PORT")
	if bp == "" {
		log.Fatal("BE_PORT env var is missing.")
	}

	url, err := url.Parse("http://localhost:" + bp)
	if err != nil {
		log.Fatalf("Error parsing backend url: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("request dump: %v", req)
		proxy.ServeHTTP(rw, req)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", fp), nil)
}
