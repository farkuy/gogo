package main

import (
	"fmt"
	"net/http"
	"os"

	transport "github.com/farkuy/gogo/cmd/internal/transport"
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
	fmt.Println("Server start on:", port)

	mux := http.NewServeMux()

	transport.RegisterRoutes(mux)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
