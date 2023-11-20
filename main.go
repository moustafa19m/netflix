package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/moustafa19m/netflix/cmd/gif"
)

const (
	port         = ":8080"
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 120 * time.Second
)

func main() {
	// get api key
	apiKey := os.Getenv("GIPHY_API_KEY")

	// create logger
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// create app
	app, err := gif.NewApp(apiKey, log)
	if err != nil {
		log.Fatal(err)
	}

	// create server
	srv := &http.Server{
		Addr:         port,
		Handler:      app,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	// start server
	log.Println("Listing for requests at http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
