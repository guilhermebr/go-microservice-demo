package main

import (
	"os"

	"github.com/guilhermebr/go-microservice-demo/v4/auth/api"
	"github.com/guilhermebr/go-microservice-demo/v4/auth/storage/mongodb"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.StandardLogger()
	log.SetLevel(logrus.DebugLevel)
	log.Infoln("Starting api...")

	endpoint := os.Getenv("AUTH_MONGODB_ENDPOINT")
	if endpoint == "" {
		endpoint = "mongodb://localhost:27017/auth"
	}

	log.Infoln("connecting to mongodb at " + endpoint)
	db, err := mongodb.New(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	authService := &api.Service{
		User: mongodb.NewUserStorage(db),
	}
	authService.SetLogger(log)

	if err := api.Start(authService); err != nil {
		log.WithError(err).Fatalln("Couldn't start api!")
	}
}
