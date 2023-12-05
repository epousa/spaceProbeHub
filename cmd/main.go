package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var url string = "https://api.nasa.gov/planetary/apod?api_key="
var method string = "GET"

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	authKey := os.Getenv("APOD_KEY")
	if authKey == "" {
		log.Fatal("AUTH_TOKEN not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url+authKey, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
