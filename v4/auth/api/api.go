package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type config struct {
	port      string
	secretKey string
}

var Config = config{}

func loadConfig() {
	Config.secretKey = os.Getenv("SECRET_KEY")
	if len(Config.secretKey) == 0 {
		log.Fatal("SECRET_KEY env var is required")
	}
	Config.port = os.Getenv("AUTH_PORT")
	if Config.port == "" {
		Config.port = "5000"
	}
}

func Start(service *Service) error {
	loadConfig()

	//Router
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", service.healthcheck).Methods("GET")
	r.HandleFunc("/user", service.createUser).Methods("POST")
	r.HandleFunc("/login", service.login).Methods("POST")

	//Negroni
	n := negroni.Classic()
	n.UseHandler(r)

	service.log.Infoln("Listen at 0.0.0.0:" + Config.port)
	http.ListenAndServe(":"+Config.port, n)

	return nil
}
