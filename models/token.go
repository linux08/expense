package models

import jwt "github.com/dgrijalva/jwt-go"

//Token model declaration
type Token struct {
	UserID uint
	*jwt.StandardClaims
}
