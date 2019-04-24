package api

import (
	"encoding/json"
	"net/http"

	"github.com/guilhermebr/go-microservice-demo/v4/auth/types"
	"github.com/sirupsen/logrus"
)

// title: create user
// path: /post
// method: POST
// responses:
//		201: created
//		400: bad request
//		500: server error
func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithFields(logrus.Fields{
		"handler": "createuser",
	})
	log.Infoln("createUser called")

	var form types.User

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Error(err)
		ErrInvalidJson.Send(w)
		return
	}

	if form.Email == "" {
		e := AlertMissingData
		e.Message = "Missing Email field"
		e.Send(w)
		return
	} else if form.Password == "" {
		e := AlertMissingData
		e.Message = "Missing Password field"
		e.Send(w)
		return
	} else if form.Role == "" {
		e := AlertMissingData
		e.Message = "Missing Role field"
		e.Send(w)
		return
	}

	//generate password hash
	form.GeneratePasswordHash(form.Password)

	if err := s.User.Create(&form); err != nil {
		log.WithField("err", err).Error("cannot create user")
		ErrInternalServer.Send(w)
		return
	}

	form.Password = ""

	Success(form, http.StatusCreated).Send(w)
}
