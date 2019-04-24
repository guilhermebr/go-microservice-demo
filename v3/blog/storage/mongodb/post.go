package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/guilhermebr/go-microservice-demo/v3/blog/types"
)

var _ types.PostStorage = &postStorage{}

const (
	postCollectionName = "post"
)

type postStorage struct {
	session *mgo.Session
}

type post struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string
	Body  string
}

func NewPostStorage(session *mgo.Session) *postStorage {
	return &postStorage{session}
}

func (s *postStorage) CloseSession() {
	s.session.Close()
}

func (s *postStorage) Create(form *types.Post) error {
	c := s.session.DB("").C(postCollectionName)
	var p post
	p.Title = form.Title
	p.Body = form.Body
	p.ID = bson.NewObjectId()
	form.ID = p.ID.Hex()
	return c.Insert(p)
}

func (s *postStorage) GetByID(id string) (*types.Post, error) {
	c := s.session.DB("").C(postCollectionName)
	var p post
	err := c.FindId(bson.ObjectIdHex(id)).One(&p)
	if err != nil {
		return nil, err
	}
	ps := types.Post{
		ID:    id,
		Title: p.Title,
		Body:  p.Body,
	}
	return &ps, err
}

func (s *postStorage) GetAll() ([]*types.Post, error) {
	c := s.session.DB("").C(postCollectionName)
	var posts []post
	err := c.Find(nil).All(&posts)
	if err != nil {
		return nil, err
	}
	appendPosts := make([]*types.Post, len(posts))
	for i, p := range posts {
		ps := types.Post{
			ID:    p.ID.Hex(),
			Title: p.Title,
			Body:  p.Body,
		}
		appendPosts[i] = &ps
	}
	return appendPosts, nil
}
