package helper_jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Verify(tokenString string) (*jwt.Token, error) {
	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(tokenString), nil
	// })
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		logrus.Error("token : ", token, err.Error())
		return nil, err
	}

	return token, nil
}
