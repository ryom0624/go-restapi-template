package main

import (
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r, err := NewApp()
	if err != nil {
		log.Fatalf("Error loading app: %s", err.Error())
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
