package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gifsearch",
		Short: "gifsearch is a way to find gifs",
		Run:   gifSearch,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg, what gif are you looking for?")
			}
			return nil
		},
	}
	rootCmd.Flags().Int("count", 1, "number of gifs to return")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gifSearch(cmd *cobra.Command, args []string) {
	random := false

	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		log.Fatalf("error getting 'count' flag: %s", err)
	}
	if count > 20 {
		log.Fatalf("only up to 20 gifs at a time please")
	}
	// If count is not specified, choose a random one out of 50
	if count == 1 {
		count = 50
		random = true
	}

	searchTerm := strings.Join(args[0:], " ")

	log.Printf("Powered By Tenor")
	log.Printf("Searching for %s...", searchTerm)

	// Send request for gifs
	u, err := url.Parse(searchURL)
	if err != nil {
		log.Fatalf("url parse error: %s", err.Error())
	}
	q := u.Query()
	q.Set("key", apiKey)
	q.Add("q", searchTerm)
	q.Add("limit", strconv.Itoa(count))
	q.Add("safesearch", "moderate")
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

	// Unmarshal response
	gifResp := TenorResponse{}
	json.Unmarshal(body, &gifResp)

	if len(gifResp.Results) == 0 {
		fmt.Printf("ಥ_ಥ  no giffy found \n")
		os.Exit(0)
	}

	if random {
		printRandom(gifResp)
	} else {
		printAll(gifResp)
	}
}

func printRandom(gifResp TenorResponse) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	result := gifResp.Results[r.Intn(len(gifResp.Results))]
	media := result.Media[0]
	url := media.Gif.URL
	body := getImage(url)

	// create tmp gif file
	file, err := os.Create("/tmp/tmpGif")
	if err != nil {
		log.Fatalf("Error creating tmp file: %s", err.Error())
	}

	// write to tmp file
	_, err = file.Write(body)
	if err != nil {
		log.Fatalf("Error writting gif to tmp file: %s", err.Error())
	}

	// print inline image
	printImage(url, file, body)
}

func printAll(gifResp TenorResponse) {
	// Get all gif urls at once but
	// print one image at a time
	var wg sync.WaitGroup
	printCh := make(chan func())

	go func(wg *sync.WaitGroup, printCh chan func()) {
		for f := range printCh {
			f()
			wg.Done()
		}
	}(&wg, printCh)

	// loop through responses
	wg.Add(len(gifResp.Results))
	for _, result := range gifResp.Results {
		media := result.Media[0]
		go func(wg *sync.WaitGroup, printCh chan func(), media Media) {
			url := media.Gif.URL
			body := getImage(url)

			// create tmp gif file
			file, err := os.Create("/tmp/tmpGif")
			if err != nil {
				log.Fatalf("Error creating tmp file: %s", err.Error())
			}

			// write to tmp file
			_, err = file.Write(body)
			if err != nil {
				log.Fatalf("Error writting gif to tmp file: %s", err.Error())
			}

			// print inline image
			printCh <- func() { printImage(url, file, body) }
		}(&wg, printCh, media)
	}

	// wait for all images to be printed
	wg.Wait()
	close(printCh)
}

func getImage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting original gif: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("response code: %s", resp.Status)
		log.Fatalf("response: %s", body)
	}
	return body
}

func printImage(url string, file *os.File, body []byte) {
	b64FileName := base64.StdEncoding.EncodeToString([]byte(file.Name()))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("%s\n\033]1337;File=name=%s;inline=1:%s\a\n", url, b64FileName, b64FileContents)
}
