package main

import (
	"github.com/sirupsen/logrus"
	"github.com/guilhermebr/go-microservice-demo/auth/api"
)

func main() {
	log := logrus.StandardLogger()
	log.Infoln("Starting api...")

	db, err := memory.New()
	if err != nil {
		log.Fatal(err)
	}

	authService := &types.Service {
		log: log,
		storage: types.Storage {
			Users: memory.NewUserStorage(db)
		}
	}

	if err := api.Start(authService); err != nil {
		log.WithError(err).Fatalln("Couldn't start api!")
	}
}
