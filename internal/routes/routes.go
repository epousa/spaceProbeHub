package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// convert JSON into GO struct -> https://mholt.github.io/json-to-go/
type apodResponse struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	var result apodResponse

	apiURL := os.Getenv("APOD") + os.Getenv("NASA_KEY") // Replace with your API endpoint
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch data from the API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read data from the API", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		log.Fatal(err)
	}

	fmt.Fprintln(w, result)
}
