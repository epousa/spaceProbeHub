package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apodResponse struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func getRequest(client *http.Client, url string) apodResponse {
	var result apodResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		log.Fatal(err)
	}

	return result
}

func main() {
	client := &http.Client{}

	apiData := getRequest(client, os.Getenv("APOD")+os.Getenv("NASA_KEY"))

	fmt.Println(apiData)
}
