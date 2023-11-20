package giphy_client

import (
	"encoding/json"
	"log"
)

// struct to map to giphy's API response
type giphyResponse struct {
	Data []struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	} `json:"data"`
}

// search query is limited to 25 results per request
func (c *GiphyClient) Search(term string, logger *log.Logger) Response {
	// setting up request to giphy
	params := map[string]string{
		"q":      term,
		"limit":  "25",
		"offset": "0",
		"rating": "g",
		"lang":   "en",
		"bundle": "messaging_non_clips",
	}

	// making the request
	logger.Printf("making request to Giphy's API for search term=%s\n", term)
	resp, err := c.makeRequest("GET", "https://api.giphy.com/v1/gifs/search", params)
	if err != nil {
		logger.Printf("can't make request: %s\n", err.Error())
		return Response{
			SearchTerm: term,
			Error:      "can't make request to Giphy's API",
		}
	}

	defer resp.Body.Close()

	// parsing the response
	var giphyResp giphyResponse
	err = json.NewDecoder(resp.Body).Decode(&giphyResp)
	if err != nil {
		logger.Printf("can't parse response: %s\n", err.Error())
		return Response{
			SearchTerm: term,
			Error:      "can't parse response",
		}
	}

	// mapping the response to the Response struct
	var gifs []Gif
	for _, v := range giphyResp.Data {
		gifs = append(gifs, Gif{
			Id:  v.Id,
			Url: v.Url,
		})
	}

	return Response{
		SearchTerm: term,
		Gifs:       gifs,
	}
}
