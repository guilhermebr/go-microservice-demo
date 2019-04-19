package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guilhermebr/go-microservice-demo/auth/types"
	"github.com/sirupsen/logrus"
)

// title: login user
// path: /auth
// method: POST
// responses:
//		200: authenticated
//		400: bad request
//		500: server error
func (s *types.Service) login(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithFields(logrus.Fields{
		"service": "auth",
		"handler": "login",
	})
	log.Infoln("login called")

	var params struct {
		Email    string
		Password string
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Error(err)
		ErrInvalidJson.Send(w)
		return
	}

	user, err := s.storage.User.GetByEmail(params.Email)
	if err != nil {
		log.Error(err)
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		AlertUserNotFound.Send(w)
		return
	}

	if !auth.CheckPasswordHash(params.Password) {
		AlertUserWrongPassword.Send(w)
		return
	}

	// generate access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp":        time.Now().Add(time.Hour * time.Duration(60)).Unix(),
		"id":         user.ID,
		"email":      user.Email,
		"profile_id": user.ProfileId,
	})

	tokenString, error := token.SignedString([]byte(SecretKey))

	if error != nil {
		e := ErrInternalServer
		e.Message = "Token not signed."
		e.Send(w)
		log.Error(error)
	}

	user.Token = "Bearer " + tokenString

	Success(auth, http.StatusOK).Send(w)
}
