package gif

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moustafa19m/netflix/pkg/giphy_client"
	"github.com/moustafa19m/netflix/pkg/web"
)

type App struct {
	http.Handler
	gifClient *giphy_client.GiphyClient
	logger    *log.Logger
}

func NewApp(gifyKey string, logger *log.Logger) (*App, error) {
	// initialize giphy client, which will be used to make requests to giphy
	gifClient, err := giphy_client.NewGiphyClient(gifyKey)
	if err != nil {
		return nil, err
	}

	// initialize router
	r := mux.NewRouter()
	a := &App{r, gifClient, logger}

	// handling 404 requests
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// handling requests
	r.HandleFunc("/query", a.searchGifs).Methods("GET")
	// ADD MORE REQUESTS BELOW

	a.Handler = r
	return a, nil
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	web.RespondError(w, http.StatusBadRequest, errors.New("not found"))
}
