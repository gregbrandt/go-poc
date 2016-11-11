package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gregbrandt/Go-POC/apihandler"
	_ "github.com/gregbrandt/Go-POC/infrastructure"
	
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
	//config := infrastructure.GetConfiguration()
	// connection := fmt.Sprintf("%s:%s@/social_network", config.DBUser, config.DBPassword)
	// db, err := sql.Open("mysql", connection)

	// if err != nil {

	// }

	// logger := infrastructure.GetLogger()
	// logger.WithFields(logrus.Fields{
	// 	"DBUser":     config.DBUser,
	// 	"DBPassword": config.DBPassword,
	// }).Info("config values")

	routes := mux.NewRouter()

	routes.HandleFunc("/api/story/create", apihandler.CreateStory)

	routes.HandleFunc("/api/user/story/{id:[0-9]+}", apihandler.GetStory)

	http.Handle("/", routes)

	log.Println("Starting redirection server")

    err := http.ListenAndServe(HTTPport, nil)
    if err != nil {
		panic(fmt.Errorf("could not start http listener: %s", err))
    }
}
