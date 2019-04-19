package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"gitlab.globoi.com/open-source/opensource-api/types"
)

type config struct {
	port          string
	sessionSecret string
}

var Config = config{}

func loadConfig() {
	Config.sessionSecret = os.Getenv("SESSION_SECRET")
	if len(Config.sessionSecret) == 0 {
		log.Fatal("SESSION_SECRET env var is required")
	}
	Config.port = os.Getenv("PORT")
	if Config.port == "" {
		Config.port = "5000"
	}
}

func Start(service types.Service) error {
	loadConfig()

	//Router
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", service.healthcheck).Methods("GET")
	r.HandleFunc("/login", service.login).Methods("POST")
	r.HandleFunc("/user", service.user).Methods("POST")

	//Negroni
	n := negroni.Classic()

	http.ListenAndServe(":"+Config.port, n)

	return nil
}
