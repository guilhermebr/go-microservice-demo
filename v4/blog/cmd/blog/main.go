package main

import (
	"os"

	"github.com/guilhermebr/go-microservice-demo/v4/blog/api"
	"github.com/guilhermebr/go-microservice-demo/v4/blog/storage/mongodb"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.StandardLogger()
	log.SetLevel(logrus.DebugLevel)
	log.Infoln("Starting api...")

	endpoint := os.Getenv("BLOG_MONGODB_ENDPOINT")
	if endpoint == "" {
		endpoint = "mongodb://localhost:27017/blog"
	}

	log.Infoln("connecting to mongodb at " + endpoint)
	db, err := mongodb.New(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	service := &api.Service{
		Post: mongodb.NewPostStorage(db),
	}
	service.SetLogger(log)

	if err := api.Start(service); err != nil {
		log.WithError(err).Fatalln("Couldn't start api!")
	}
}
