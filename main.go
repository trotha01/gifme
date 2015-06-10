package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type giphyImage struct {
	URL string `json:"url"`
}

type giphyImages struct {
	Original giphyImage `json:"original"`
}

type giphyData struct {
	Images giphyImages `json:"images"`
}

type giphy struct {
	Data []giphyData `json:"data"`
}

type MyJsonName struct {
	Data []struct {
		Images struct {
			Original struct {
				URL string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}

const searchURL = "http://api.giphy.com/v1/gifs/search"
const apiKey = "dc6zaTOxFJmzC"

func main() {
	searchTerm := "doggie"
	reqURL := fmt.Sprintf("%s?q=%s&api_key=%s", searchURL, searchTerm, apiKey)
	log.Printf("req: %s", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatalf("Error getting gif: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("response code: %s", resp.Status)
		log.Fatalf("response: %s", body)
	}

	// gifResp := MyJsonName{}
	gifResp := giphy{}
	json.Unmarshal(body, &gifResp)
	log.Printf("Gif: %s", body)
	log.Printf("Gif: %+v", gifResp)
}
