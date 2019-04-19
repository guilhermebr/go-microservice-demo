package types

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        string
	Email     string
	Password  string
	Role      string
	ProfileId string
	Token     string `json:"access_token,omitempty"`
}

type UserStorage interface {
	Create(*User) (*User, error)
	GetByEmail(string) (*User, error)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
