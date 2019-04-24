package memory

import "github.com/guilhermebr/go-microservice-demo/v2/blog/types"

type DB struct {
	posts map[string]*types.Post
}

func New() *DB {
	return &DB{
		posts: make(map[string]*types.Post),
	}
}
