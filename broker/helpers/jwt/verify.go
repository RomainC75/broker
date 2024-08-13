package helper_jwt

import (
	"fmt"
	"shared/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Verify(tokenString string) error {
	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(tokenString), nil
	// })
	fmt.Println("========= Verify")
	token, parts, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	utils.PrettyDisplay("parts :", parts)

	if err != nil {
		logrus.Error("token : ", token, err.Error())
		return err
	}
	utils.PrettyDisplay("TOKEN : ", *token)

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
