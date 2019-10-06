package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", Router))
}
