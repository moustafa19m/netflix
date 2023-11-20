package giphy_client

import (
	"fmt"
	"net/http"
)

func (c *GiphyClient) makeRequest(method string, url string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("giphy api returned status code %d", resp.StatusCode)
	}
	return resp, err
}
