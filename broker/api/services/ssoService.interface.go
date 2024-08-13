package services

type SsoServiceInterface interface {
	GetPublicKeys(token string) ([]JWKS, error)
	IsKidInPublicKeys(token string, keys []JWKS)
}
