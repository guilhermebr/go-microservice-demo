package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermebr/go-microservice-demo/v3/blog/types"
	"github.com/sirupsen/logrus"
)

// title: create post
// path: /post
// method: POST
// responses:
//		201: created
//		400: bad request
//		500: server error
func (s *Service) createPost(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithFields(logrus.Fields{
		"handler": "createPost",
	})
	log.Infoln("createPost called")

	var form types.Post

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Error(err)
		ErrInvalidJson.Send(w)
		return
	}

	if form.Title == "" {
		e := AlertMissingData
		e.Message = "Missing Title field"
		e.Send(w)
		return
	} else if form.Body == "" {
		e := AlertMissingData
		e.Message = "Missing Body field"
		e.Send(w)
		return
	}

	if err := s.Post.Create(&form); err != nil {
		log.WithField("err", err).Error("cannot insert form")
		ErrInternalServer.Send(w)
		return
	}

	updatedForm, err := s.Post.GetByID(form.ID)
	if err != nil {
		log.WithField("err", err).Error("cannot get post")
		updatedForm = &form
	}
	Success(updatedForm, http.StatusCreated).Send(w)
}

// title: list posts
// path: /post
// method: GET
// responses:
//		200: ok
//		400: bad request
//		500: server error
func (s *Service) listPosts(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithFields(logrus.Fields{
		"handler": "listPosts",
	})
	log.Infoln("listPosts called")

	posts, err := s.Post.GetAll()
	if err != nil {
		log.WithField("err", err).Error("cannot get posts")
		ErrInternalServer.Send(w)
		return
	}

	Success(posts, http.StatusOK).Send(w)
}

// title: get post
// path: /post/{id}
// method: GET
// responses:
//		200: ok
//		400: bad request
//		500: server error
func (s *Service) getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log := s.log.WithFields(logrus.Fields{
		"handler": "getPost",
		"id":      id,
	})
	log.Infoln("called")

	post, err := s.Post.GetByID(id)
	if err != nil {
		log.WithField("err", err).Error("cannot get post")
		ErrInternalServer.Send(w)
		return
	}

	Success(post, http.StatusOK).Send(w)
}
