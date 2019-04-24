package api

import (
	"github.com/guilhermebr/go-microservice-demo/v4/blog/types"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log  *logrus.Logger
	Post types.PostStorage
}

func (s *Service) SetLogger(log *logrus.Logger) {
	s.log = log
}
