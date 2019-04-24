package mongodb

import (
	"os"

	"github.com/globalsign/mgo"
)

func New() (*mgo.Session, error) {
	endpoint := os.Getenv("MONGODB_ENDPOINT")
	if endpoint == "" {
		endpoint = "mongodb://localhost:27017/blog"
	}

	session, err := mgo.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return session, nil
}
