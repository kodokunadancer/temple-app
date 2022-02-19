package main

import (
	"fmt"
	"log"
	"net/http"
)

func routing() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", logMiddleware(http.HandlerFunc(hello)))
	return mux
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		next.ServeHTTP(w, r)
		log.Println("after")
	}
}

// レスポンスで"Hello World"と返すAPI
func hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
