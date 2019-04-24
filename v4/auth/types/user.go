package types

import "golang.org/x/crypto/bcrypt"

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type UserRole string

type User struct {
	Email    string
	Password string
	Role     UserRole
}

type UserStorage interface {
	Create(*User) error
	GetByEmail(string) (*User, error)
}

func (user *User) GeneratePasswordHash(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	user.Password = string(bytes)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
