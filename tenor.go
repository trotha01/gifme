package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const tenorSearchURL = "http://api.tenor.com/v1/search"
const tenorAPIKey = "CH97GP0E42LR"

// TenorResponse is the gif response from tenor.com
type TenorResponse struct {
	Results []GifResult `json:"results"`
}

// GifResult is an individual gif result
type GifResult struct {
	URL   string  `json:"url"`
	Media []Media `json:"media"`
}

// Media is a return value in a tenor response
type Media struct {
	Gif struct {
		URL string `json:"url"`
	} `json:"gif"`
}

func tenorGetGifs(searchTerm string, count int) (TenorResponse, error) {
	gifResp := TenorResponse{}

	// Create Tenor http request
	u, err := url.Parse(tenorSearchURL)
	if err != nil {
		log.Fatalf("url parse error: %s", err.Error())
	}
	q := u.Query()
	q.Set("key", tenorAPIKey)
	q.Add("q", searchTerm)
	// Always get 50 gifs and pick randomly from results
	q.Add("limit", "50")
	q.Add("safesearch", "moderate")
	u.RawQuery = q.Encode()

	// Send request for gifs
	resp, err := http.Get(u.String())
	if err != nil {
		return gifResp, fmt.Errorf("Tenor http request err: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return gifResp, fmt.Errorf("Tenor http request non 200 response: %s (%s)", resp.Status, body)
	}

	// Unmarshal response
	err = json.Unmarshal(body, &gifResp)
	if err != nil {
		return gifResp, fmt.Errorf("Error unmarshalling Tenor resposne: %s", body)
	}

	if len(gifResp.Results) == 0 {
		return gifResp, fmt.Errorf("ಥ_ಥ  no giffy found")
	}

	return gifResp, nil
}

func tenorGetImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error getting gif from Tenor: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Tenor gif http request non 200 response: %s (%s)", resp.Status, body)
	}
	return body, nil
}
