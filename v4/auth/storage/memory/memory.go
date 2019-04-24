package memory

import "github.com/guilhermebr/go-microservice-demo/types"

type DB struct {
	users map[string]*types.User
}

func New() *DB {
	return &DB{
		users: make(map[string]*types.User),
	}
}
