package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Start(service *Service) error {
	//Router
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", service.healthcheck).Methods("GET")
	r.HandleFunc("/post", service.createPost).Methods("POST")
	r.HandleFunc("/post", service.listPosts).Methods("GET")
	r.HandleFunc("/post/{id}", service.getPost).Methods("GET")

	//Negroni
	n := negroni.Classic()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.ListenAndServe(":"+port, n)

	return nil
}
