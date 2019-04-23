package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	log := logrus.StandardLogger()
	log.Infoln("Starting api...")

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheck).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")

	//Negroni
	n := negroni.Classic()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.ListenAndServe(":"+port, n)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
