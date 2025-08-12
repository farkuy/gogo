package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {

	loadDotEnv()
	port, exits := os.LookupEnv("PORT")

	if !exits {
		fmt.Println("The port was not found and start on 8000")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
