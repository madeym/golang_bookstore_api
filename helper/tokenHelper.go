package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SigningDetails struct {
	ID    int
	Name  string
	Email string
	jwt.StandardClaims
}

var UserData SigningDetails

func ValidateToken(signedToken string) (claims *SigningDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SigningDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SigningDetails)
	if !ok {
		msg = fmt.Sprintf("token invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	UserData.ID = claims.ID
	UserData.Name = claims.Name
	UserData.Email = claims.Email

	return claims, msg
}
