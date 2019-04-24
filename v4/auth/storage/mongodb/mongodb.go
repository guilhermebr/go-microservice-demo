package mongodb

import (
	"github.com/globalsign/mgo"
)

func New(endpoint string) (*mgo.Session, error) {
	session, err := mgo.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return session, nil
}
