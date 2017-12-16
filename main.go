package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
			count, err := cmd.Flags().GetInt("count")
			if err != nil {
				log.Fatalf("error getting 'count' flag: %s", err)
			}
			if count < 0 {
				return fmt.Errorf("'count' flag must not be negative, got :%d", count)
			}
			if count > 20 {
				log.Fatalf("only up to 20 gifs at a time please, got %d", count)
			}
			return nil
		},
	}
	rootCmd.Flags().IntP("count", "c", 1, "number of gifs to return")
	rootCmd.Flags().StringP("engine", "e", "", "gif engine to use 'giphy' or 'tenor'. If not specified Tenor is searched first and Gifme if there is an error from Tenor")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gifSearch(cmd *cobra.Command, args []string) {
	searchTerm := strings.Join(args[0:], " ")

	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		log.Fatalf("error getting 'count' flag: %s", err)
	}

	engine, err := cmd.Flags().GetString("engine")
	if err != nil {
		log.Fatalf("error getting 'engine' flag: %s", err)
	}

	switch {
	case engine == "giphy":
		err := giffyGifSearch(searchTerm, count)
		if err != nil {
			log.Fatal("Error getting gif from Giphy:", err)
		}
	case engine == "tenor":
		err = tenorGifSearch(searchTerm, count)
		if err != nil {
			log.Fatal("Error getting gif from Tenor:", err)
		}
	default:
		err = tenorGifSearch(searchTerm, count)
		if err != nil {
			log.Println("Error getting gif from Tenor:", err)
			// If error, try Giffy next
			err = giffyGifSearch(searchTerm, count)
			if err != nil {
				log.Println("Error getting gif from Giphy:", err)
			}
		}
	}

}

func giffyGifSearch(searchTerm string, count int) error {
	log.Printf("Powered By Giphy")
	log.Printf("Searching for %s...", searchTerm)

	gifResp, err := giffyGetGifs(searchTerm, count)
	if err != nil {
		return err
	}

	gifCount := len(gifResp.Data)
	if gifCount == 0 {
		return fmt.Errorf("ಥ_ಥ  no giffy found")
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(gifResp.Data))

	for i := 0; i < count; i++ {
		gifChoice := gifResp.Data[perm[i]]

		// get chosen gif
		resp, err := http.Get(gifChoice.Images.Original.URL)
		if err != nil {
			return fmt.Errorf("Error getting original gif: %s", err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			log.Printf("response code: %s", resp.Status)
			log.Fatalf("response: %s", body)
		}

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
		printImage(gifChoice.Images.Original.URL, file, body)
	}
	return nil
}

func tenorGifSearch(searchTerm string, count int) error {
	log.Printf("Powered By Tenor")
	log.Printf("Searching for %s...", searchTerm)

	gifResp, err := tenorGetGifs(searchTerm, count)
	if err != nil {
		return err
	}

	return printCount(gifResp, count)
}

func printCount(gifResp TenorResponse, count int) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(gifResp.Results))

	for i := 0; i < count; i++ {
		// get the next random result
		result := gifResp.Results[perm[i]]
		media := result.Media[0]
		url := media.Gif.URL
		body, err := tenorGetImage(url)
		if err != nil {
			return err
		}

		// create tmp gif file
		file, err := os.Create("/tmp/tmpGif")
		if err != nil {
			return fmt.Errorf("Error creating tmp file: %s", err.Error())
		}

		// write to tmp file
		_, err = file.Write(body)
		if err != nil {
			return fmt.Errorf("Error writting gif to tmp file: %s", err.Error())
		}

		// print inline image
		printImage(url, file, body)
	}
	return nil
}

func printImage(url string, file *os.File, body []byte) {
	b64FileName := base64.StdEncoding.EncodeToString([]byte(file.Name()))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("%s\n\033]1337;File=name=%s;inline=1:%s\a\n", url, b64FileName, b64FileContents)
}
