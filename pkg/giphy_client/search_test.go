package giphy_client

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func TestSearch(t *testing.T) {
	t.Run("Giphy API Positive", func(t *testing.T) {
		defer gock.Off()

		client := GiphyClient{
			apiKey: "test",
			client: &http.Client{},
		}

		term := "test"

		testData, err := os.Open("test_data/giphy_test_positive.json")
		if err != nil {
			t.Errorf("error opening testdata: %v", err)
		}
		defer testData.Close()

		mockResponse := testData

		gock.New("https://api.giphy.com").
			Get("/v1/gifs/search").
			MatchParams(map[string]string{
				"api_key": client.apiKey,
				"q":       term,
				"limit":   "25",
				"offset":  "0",
				"rating":  "g",
				"lang":    "en",
				"bundle":  "messaging_non_clips",
			}).
			Reply(200).
			Body(mockResponse)

		res := client.Search(term, logger)

		require.Equal(t, 25, len(res.Gifs))
		require.Equal(t, "test", res.SearchTerm)
		require.Equal(t, "", res.Error)
		for _, v := range res.Gifs {
			require.NotEqual(t, "", v.Id)
			require.NotEqual(t, "", v.Url)
		}
	})

	t.Run("Giphy API Negative 500", func(t *testing.T) {
		defer gock.Off()

		client := GiphyClient{
			apiKey: "test",
			client: &http.Client{},
		}

		term := "test"

		gock.New("https://api.giphy.com").
			Get("/v1/gifs/search").
			MatchParams(map[string]string{
				"api_key": client.apiKey,
				"q":       term,
				"limit":   "25",
				"offset":  "0",
				"rating":  "g",
				"lang":    "en",
				"bundle":  "messaging_non_clips",
			}).
			Reply(500)

		res := client.Search(term, logger)

		require.Equal(t, 0, len(res.Gifs))
		require.Equal(t, "test", res.SearchTerm)
		require.Equal(t, "can't make request to Giphy's API", res.Error)
		for _, v := range res.Gifs {
			require.Equal(t, "", v.Id)
			require.Equal(t, "", v.Url)
		}
	})

	t.Run("Giphy API Negative Wrong Json Format", func(t *testing.T) {
		defer gock.Off()

		client := GiphyClient{
			apiKey: "test",
			client: &http.Client{},
		}

		term := "test"

		testData, err := os.Open("test_data/giphy_test_negative-json.txt")
		if err != nil {
			t.Errorf("error opening testdata: %v", err)
		}
		defer testData.Close()

		mockResponse := testData

		gock.New("https://api.giphy.com").
			Get("/v1/gifs/search").
			MatchParams(map[string]string{
				"api_key": client.apiKey,
				"q":       term,
				"limit":   "25",
				"offset":  "0",
				"rating":  "g",
				"lang":    "en",
				"bundle":  "messaging_non_clips",
			}).
			Reply(200).
			Body(mockResponse)

		res := client.Search(term, logger)

		require.Equal(t, 0, len(res.Gifs))
		require.Equal(t, "test", res.SearchTerm)
		require.Equal(t, "can't parse response", res.Error)
		for _, v := range res.Gifs {
			require.Equal(t, "", v.Id)
			require.Equal(t, "", v.Url)
		}
	})
}
