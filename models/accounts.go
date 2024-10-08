package models

import "github.com/dgrijalva/jwt-go"

// JWT access rights structure.
type Token struct {
	UserId uint
	jwt.StandardClaims
}
