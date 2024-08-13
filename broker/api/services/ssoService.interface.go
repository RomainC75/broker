package services

import "github.com/golang-jwt/jwt/v5"

type SsoServiceInterface interface {
	ExtractTokenClaims(tokenString string) (jwt.Claims, error)
}
