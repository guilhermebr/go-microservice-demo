package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

type AuthUser struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func GetUser(r *http.Request) AuthUser {
	var authUser AuthUser

	decoded := context.Get(r, "decoded")

	mapstructure.Decode(decoded.(jwt.MapClaims), &authUser)

	return authUser
}

func ValidateToken(secretKey string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader == "" {
			e := ErrForbidden
			e.Message = "An authorization header is required"
			e.Send(w)
			return
		}

		bearerToken := strings.Split(authorizationHeader, "Bearer ")
		if len(bearerToken) != 2 {
			e := ErrForbidden
			e.Message = "Invalid token"
			e.Send(w)
			return
		}

		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secretKey), nil
		})

		if !token.Valid {
			e := ErrUnauthorized
			e.Message = "Invalid authorization token"
			e.Send(w)
			return
		}

		context.Set(r, "decoded", token.Claims)
		next(w, r)
	})
}
