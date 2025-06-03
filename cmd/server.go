package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root endpoint\n")
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env file found or failed to load. Using default config.")
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	port := os.Getenv("API_PORT")

	fmt.Println("Server starting on " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
