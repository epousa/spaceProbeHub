package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://api.nasa.gov/planetary/apod?api_key=<api_key>"
var method string = "GET"

func main() {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
