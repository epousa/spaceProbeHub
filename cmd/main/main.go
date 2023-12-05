package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/epousa/spaceProbeHub/internal/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
