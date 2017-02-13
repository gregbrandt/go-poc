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

	storyRoot := rootURL + "/story1" 
	routes.HandleFunc(storyRoot + "/", apihandler.CreateStory).
			Methods("POST")
	routes.HandleFunc(storyRoot + "/{id:[0-9]+}", apihandler.GetStory).
			Methods("GET")
	routes.PathPrefix("/").Handler(http.FileServer(http.Dir("www/dist")))
	// routes.Handle(`/{rest:[a-zA-Z0-9=\-\/]+}`,http.FileServer(http.Dir("www/dist")))
	// r.Handle("/js", http.FileServer(http.Dir("../sitelocation/js")))
    // r.Handle("/css", http.FileServer(http.Dir("../sitelocation/css")))
	// routes.Handle(`/{[a-zA-Z0-9=\-]+\.js$}`,http.FileServer(http.Dir("www/dist")))
	// routes.Handle(`/{[a-zA-Z0-9=\-]+\.css$}`,http.FileServer(http.Dir("www/dist")))
	// routes.NotFoundHandler = http.HandlerFunc(notFound) 
	// routes.Handle("/*",http.FileServer(http.File("www/dist/index.html")))
	// routes.PathPrefix("/").Handler(http.FileServer(http.Dir("www/dist")))

	
	http.Handle("/", routes)
	
}


// func notFound(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.RawPath);
//     http.ServeFile(w, r, "www/dist/index.html")
// }