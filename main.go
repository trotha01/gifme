package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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

const searchURL = "http://api.giphy.com/v1/gifs/search"
const apiKey = "dc6zaTOxFJmzC"

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Please enter a gif to search for")
	}
	searchTerm := strings.Join(os.Args[1:], " ")

	// create tmp gif file
	file, err := os.Create("/tmp/tmpGif")
	if err != nil {
		log.Fatalf("Error creating tmp file: %s", err.Error())
	}

	log.Printf("Powered by Giphy")
	log.Printf("Searching for %s...", searchTerm)

	u, err := url.Parse(searchURL)
	if err != nil {
		log.Fatalf("url parse error: %s", err.Error())
	}
	q := u.Query()
	q.Set("api_key", apiKey)
	q.Add("q", searchTerm)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("Error getting gifs: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("response code: %s", resp.Status)
		log.Fatalf("response: %s", body)
	}

	// unmarshal response
	gifResp := giphy{}
	json.Unmarshal(body, &gifResp)

	// randomly choose a gif from the response
	r := rand.New(rand.NewSource(time.Now().Unix()))
	gifCount := len(gifResp.Data)
	gifChoice := gifResp.Data[r.Intn(gifCount)]

	// get original gif
	resp, err = http.Get(gifChoice.Images.Original.URL)
	if err != nil {
		log.Fatalf("Error getting original gif: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("response code: %s", resp.Status)
		log.Fatalf("response: %s", body)
	}

	// write to tmp file
	_, err = file.Write(body)
	if err != nil {
		log.Fatalf("Error writting gif to tmp file: %s", err.Error())
	}

	// print inline image
	b64FileName := base64.StdEncoding.EncodeToString([]byte(file.Name()))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("\033]1337;File=name=%s;inline=1:%s\a\n", b64FileName, b64FileContents)
}
