package user

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	// ID 用户id
	ID string `json:"id"`
	jwt.RegisteredClaims
}
