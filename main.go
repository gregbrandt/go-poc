package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gregbrandt/Go-POC/apihandler"
	_ "github.com/gregbrandt/Go-POC/infrastructure"
	_ "github.com/gregbrandt/Go-POC/readmodel"
	
)

const (
	serverName   = "localhost"
	SSLport      = ":443"
	HTTPport     = ":8080"
	SSLprotocol  = "https://"
	HTTPprotocol = "http://"
	randomLength = 16
)

func main() {	
	registerRoutes()
	log.Println("Starting redirection server")

    err := http.ListenAndServe(HTTPport, nil)
    if err != nil {
		panic(fmt.Errorf("could not start http listener: %s", err))
    }
}

func registerRoutes(){
	rootURL := "/api"
	routes := mux.NewRouter()

	storyRoot := rootURL + "/story" 
	routes.HandleFunc(storyRoot + "/", apihandler.CreateStory).
			Methods("POST")
	routes.HandleFunc(storyRoot + "/{id:[0-9]+}", apihandler.GetStory).
			Methods("GET")

	http.Handle("/", routes)
}