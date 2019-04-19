package types

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/bsprojectdev/backend/auth/types"
)

type Service struct {
	log     *logrus.Logger
	storage Storage
}

type Storage struct {
	User types.UserStorage
}
