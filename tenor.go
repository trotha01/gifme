package main

const searchURL = "http://api.tenor.com/v1/search"
const apiKey = "CH97GP0E42LR"

// TenorResponse is the gif response from tenor.com
type TenorResponse struct {
	Weburl  string `json:"weburl"`
	Results []struct {
		Hascaption bool          `json:"hascaption,omitempty"`
		Tags       []interface{} `json:"tags"`
		URL        string        `json:"url"`
		Media      []Media       `json:"media"`
		Created    float64       `json:"created"`
		Shares     int           `json:"shares"`
		Itemurl    string        `json:"itemurl"`
		Composite  interface{}   `json:"composite"`
		Hasaudio   bool          `json:"hasaudio"`
		Title      string        `json:"title"`
		ID         string        `json:"id"`
	} `json:"results"`
	Next string `json:"next"`
}

// Media is a return value in a tenor response
type Media struct {
	Nanomp4 struct {
		URL      string  `json:"url"`
		Dims     []int   `json:"dims"`
		Duration float64 `json:"duration"`
		Preview  string  `json:"preview"`
		Size     int     `json:"size"`
	} `json:"nanomp4"`
	Nanowebm struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"nanowebm"`
	Tinygif struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"tinygif"`
	Tinymp4 struct {
		URL      string  `json:"url"`
		Dims     []int   `json:"dims"`
		Duration float64 `json:"duration"`
		Preview  string  `json:"preview"`
		Size     int     `json:"size"`
	} `json:"tinymp4"`
	Tinywebm struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"tinywebm"`
	Webm struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"webm"`
	Gif struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"gif"`
	Mp4 struct {
		URL      string  `json:"url"`
		Dims     []int   `json:"dims"`
		Duration float64 `json:"duration"`
		Preview  string  `json:"preview"`
		Size     int     `json:"size"`
	} `json:"mp4"`
	Loopedmp4 struct {
		URL      string  `json:"url"`
		Dims     []int   `json:"dims"`
		Duration float64 `json:"duration"`
		Preview  string  `json:"preview"`
	} `json:"loopedmp4"`
	Mediumgif struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"mediumgif"`
	Nanogif struct {
		URL     string `json:"url"`
		Dims    []int  `json:"dims"`
		Preview string `json:"preview"`
		Size    int    `json:"size"`
	} `json:"nanogif"`
}
