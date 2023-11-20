package gif

import (
	"errors"
	"net/http"
	"sync"

	"github.com/moustafa19m/netflix/pkg/giphy_client"
	"github.com/moustafa19m/netflix/pkg/stringset"
	"github.com/moustafa19m/netflix/pkg/web"
	"github.com/oklog/ulid/v2"
)

func (a *App) searchGifs(w http.ResponseWriter, r *http.Request) {
	// each request will have a unique id
	requestId := ulid.Make()
	a.logger.SetPrefix("INFO: [" + requestId.String() + "] ")
	a.logger.Println("Received New Request")

	// capture multiple params all with key "searchTerm"
	searchTerms := r.URL.Query()["searchTerm"]
	if len(searchTerms) == 0 {
		a.logger.Println("Processed Aborted")
		web.RespondError(w, http.StatusBadRequest, errors.New("missing search terms: please provide at least one search term"))
		return
	}

	// create a set of search terms to eliminate duplicates
	termsSet := stringset.NewSet(searchTerms...)

	// max number of requests of beta per day is 1000, therefore I added this limit
	if len(termsSet.Elements()) > 50 {
		a.logger.Println("Processed Aborted")
		web.RespondError(w, http.StatusBadRequest, errors.New("too many search params: max allowed is 50 params per request"))
		return
	}
	a.logger.Printf("[%s] Searching for gifs with terms: %s\n", requestId, termsSet.ToString())

	// for each search param run a go routine that will make a request to giphy
	// save the results in a slice
	// wait for all go routines to finish
	results := []giphy_client.Response{}
	var wg sync.WaitGroup
	for _, v := range termsSet.Elements() {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			// get gifs from giphy
			gifs := a.gifClient.Search(v, a.logger)
			// append to results
			results = append(results, gifs)
		}(v)
	}
	wg.Wait()

	a.logger.Println("Processed Request")
	// return results
	web.RespondSuccess(w, results)
}
