package pkg

import "github.com/golang-jwt/jwt/v5"

type UserToken struct {
	jwt.RegisteredClaims
	UserId string
	Email  string
}
