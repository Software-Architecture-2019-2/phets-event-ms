package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, Router))
}
