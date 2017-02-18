package main

import (
	"fmt"
	"strings"
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
	//routes.PathPrefix("/").Handler(http.FileServer(http.Dir("www/dist")))
	// routes.Handle(`/{rest:[a-zA-Z0-9=\-\/]+}`,http.FileServer(http.Dir("www/dist")))
	//routes.Handle(`/vendor.2a4c3e544c08ac8adc46.js`, http.FileServer(http.Dir("www/dist")))
	//routes.Handle(`/app.2a4c3e544c08ac8adc46.js`, http.FileServer(http.Dir("www/dist")))
     //routes.Handle(`*\.css`,http.FileServer(http.Dir("www/dist")))
//	 routes.HandleFunc(`\/*[a-zA-Z0-9=\/\-\.]+\.js$`,jscsHandler)
//	 routes.HandleFunc(`\/*[a-zA-Z0-9=\/\-\.]+\.css$}`,jscsHandler)
	routes.NotFoundHandler = http.HandlerFunc(notFound) 
	// routes.Handle("/*",http.FileServer(http.File("www/dist/index.html")))
	// routes.PathPrefix("/").Handler(http.FileServer(http.Dir("www/dist")))

	
	http.Handle("/", routes)
	
}


func notFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.RawPath);
	
	if(strings.HasSuffix(r.URL.Path,".map")){
    	http.ServeFile(w, r, "www/dist" + r.URL.Path)	
	}


	if(strings.HasSuffix(r.URL.Path,".js")){
    	http.ServeFile(w, r, "www/dist" + r.URL.Path)	
	}

    if(strings.HasSuffix(r.URL.Path,".css")){
    	http.ServeFile(w, r, "www/dist" + r.URL.Path)	
	}

	http.ServeFile(w, r, "www/dist/index.html")
}

