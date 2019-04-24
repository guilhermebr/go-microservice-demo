package api

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// title: login user
// path: /login
// method: POST
// responses:
//		200: authenticated
//		400: bad request
//		500: server error
func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithFields(logrus.Fields{
		"handler": "login",
	})
	log.Infoln("login called")

	var form struct {
		Email    string
		Password string
	}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Error(err)
		ErrInvalidJson.Send(w)
		return
	}

	user, err := s.User.GetByEmail(form.Email)
	if err != nil {
		log.WithField("err", err).Error("cannot authenticate user")
		ErrInternalServer.Send(w)
		return
	}

	if user.Email == "" {
		AlertUserNotFound.Send(w)
		return
	}

	if !user.CheckPasswordHash(form.Password) {
		AlertUserWrongPassword.Send(w)
		return
	}

	// generate access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * time.Duration(60)).Unix(),
		"email": user.Email,
		"role":  user.Role,
		//	"profile_id": user.ProfileId,
	})

	tokenString, error := token.SignedString([]byte(Config.secretKey))

	if error != nil {
		e := ErrInternalServer
		e.Message = "Token not signed."
		e.Send(w)
		log.Error(error)
	}

	Success(tokenString, http.StatusOK).Send(w)
}
