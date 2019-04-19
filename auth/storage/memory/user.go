package memory

import (
	"errors"

	"github.com/guilhermebr/go-microservice-demo/types"
)

type UserStorage struct {
	db *DB
}

func NewUserStorage(db *DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) Create(user *types.User) (*types.User, error) {
	if _, ok := s.db.users[user.Email]; ok {
		return errors.New("user already exist")
	}
	s.db.users[user.Email] = user
	return nil
}

func (s *UserStorage) GetByEmail(email string) (*types.User, error) {
	if user, ok := s.db.users[email]; ok {
		u := *user
		return &u, nil
	}
	return nil, errors.New("user not found")
}
