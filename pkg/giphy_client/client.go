package giphy_client

import (
	"errors"
	"net/http"
)

type GiphyClient struct {
	apiKey string
	client *http.Client
}

type Gif struct {
	Id  string `json:"gif_id"`
	Url string `json:"url"`
}

// search query is limited to 25 results per request
// Response returned from client to match project requirements
type Response struct {
	SearchTerm string `json:"search_term"`
	Gifs       []Gif  `json:"gifs"`
	Error      string `json:"error,omitempty"`
}

// creating new client
func NewGiphyClient(key string) (*GiphyClient, error) {
	if key == "" {
		return nil, errors.New("api key is required, given an empty string")
	}
	return &GiphyClient{key, &http.Client{}}, nil
}
