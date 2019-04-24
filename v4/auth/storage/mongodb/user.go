package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/guilhermebr/go-microservice-demo/v4/auth/types"
)

var _ types.UserStorage = &userStorage{}

const (
	userCollectionName = "user"
)

type userStorage struct {
	session *mgo.Session
}

type user struct {
	Email    string `bson:"_id"`
	Password string `json:"-"`
	Role     types.UserRole
	//	ProfileID string `json:"profile_id"`
}

func NewUserStorage(session *mgo.Session) *userStorage {
	return &userStorage{session}
}

func (s *userStorage) CloseSession() {
	s.session.Close()
}

func (s *userStorage) Create(form *types.User) error {
	c := s.session.DB("").C(userCollectionName)
	return c.Insert(user(*form))
}

func (s *userStorage) GetByEmail(email string) (*types.User, error) {
	c := s.session.DB("").C(userCollectionName)
	var u user
	err := c.FindId(email).One(&u)
	if err != nil {
		return nil, err
	}
	us := types.User(u)
	return &us, err
}
