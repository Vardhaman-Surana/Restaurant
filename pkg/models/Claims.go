package models

import "github.com/dgrijalva/jwt-go"


type Claims struct{
	UserName string
	jwt.StandardClaims
}
