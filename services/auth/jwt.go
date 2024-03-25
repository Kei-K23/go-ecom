package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserId  int
	Expires int64
	jwt.RegisteredClaims
}

func CreateJWT(secret []byte, userID int) (string, error) {

	expires := time.Second * time.Duration(3600*24*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		UserId:  userID,
		Expires: time.Now().Add(expires).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
