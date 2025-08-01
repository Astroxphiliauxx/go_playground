package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://api.restful-api.dev/objects"

func main() {
	sendPostRequest(url)
}

func sendPostRequest(url string) {
	//fake payload data

	requestBody := strings.NewReader(`
	   {

	    "course": "Go Programming",
	    "duration": "4 weeks",
	    "instructor": "John Doe",
	    "level": "Intermediate"
	    }
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
