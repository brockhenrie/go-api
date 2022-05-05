package util

import (
	"time"
	"github.com/dgrijalva/jwt-go/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string)(string, error){
	claims := jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	return signedToken, err
}

func ParseJwt(cookie string)(string, error){
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid{
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}