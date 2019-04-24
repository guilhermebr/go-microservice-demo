package main

import (
	"github.com/guilhermebr/go-microservice-demo/v3/blog/api"
	"github.com/guilhermebr/go-microservice-demo/v3/blog/storage/memory"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.StandardLogger()
	log.Infoln("Starting api...")

	db := memory.New()

	service := &api.Service{
		Post: memory.NewPostStorage(db),
	}
	service.SetLogger(log)

	if err := api.Start(service); err != nil {
		log.WithError(err).Fatalln("Couldn't start api!")
	}
}
