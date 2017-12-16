package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const giphySearchURL = "http://api.giphy.com/v1/gifs/search"
const giphyAPIKey = "dc6zaTOxFJmzC"

type giphy struct {
	Data []giphyData `json:"data"`
}

type giphyData struct {
	Images struct {
		Original struct {
			URL string `json:"url"`
		} `json:"original"`
	} `json:"images"`
}

func giffyGetGifs(searchTerm string, count int) (giphy, error) {
	gifResp := giphy{}

	// Create Giphy http request
	u, err := url.Parse(giphySearchURL)
	if err != nil {
		return gifResp, fmt.Errorf("url parse error: %s", err.Error())
	}
	q := u.Query()
	q.Set("api_key", giphyAPIKey)
	q.Add("q", searchTerm)
	u.RawQuery = q.Encode()

	// Send request for gifs
	resp, err := http.Get(u.String())
	if err != nil {
		return gifResp, fmt.Errorf("Giphy http request err: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return gifResp, fmt.Errorf("Giphy http request non 200 response: %s (%s)", resp.Status, body)
	}

	// Unmarshal response
	err = json.Unmarshal(body, &gifResp)
	if err != nil {
		return gifResp, fmt.Errorf("Error unmarshalling Giphy resposne: %s", body)
	}

	return gifResp, nil
}
