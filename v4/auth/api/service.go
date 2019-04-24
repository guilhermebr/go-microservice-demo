package api

import (
	"github.com/guilhermebr/go-microservice-demo/v4/auth/types"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log  *logrus.Logger
	User types.UserStorage
}

func (s *Service) SetLogger(log *logrus.Logger) {
	s.log = log
}
