package memory

import (
	"errors"
	"strconv"

	"github.com/guilhermebr/go-microservice-demo/v3/blog/types"
)

var _ types.PostStorage = &postStorage{}

type postStorage struct {
	db           *DB
	postSequence int
}

func NewPostStorage(db *DB) *postStorage {
	return &postStorage{db: db}
}

func (s *postStorage) nextID() string {
	s.postSequence += 1
	return string(strconv.Itoa(s.postSequence))
}

func (s *postStorage) Create(post *types.Post) error {
	post.ID = s.nextID()
	s.db.posts[post.ID] = post
	return nil
}

func (s *postStorage) GetByID(id string) (*types.Post, error) {
	if post, ok := s.db.posts[id]; ok {
		p := *post
		return &p, nil
	}
	return nil, errors.New("post not found")
}

func (s *postStorage) GetAll() ([]*types.Post, error) {
	var posts []*types.Post
	for _, p := range s.db.posts {
		np := *p
		posts = append(posts, &np)
	}
	return posts, nil
}
