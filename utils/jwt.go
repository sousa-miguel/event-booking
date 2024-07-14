package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(GetDotEnvVariable("SUPER_SECRET")))
}
