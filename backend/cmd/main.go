package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routing()
	log.Fatal(http.ListenAndServe(":8081", mux))
}
