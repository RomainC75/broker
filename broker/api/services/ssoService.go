package services

import (
	dto_response "broker/api/dto/response"
	helper_jwt "broker/helpers/jwt"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shared/config"
	redis_repo "shared/repositories/redis"
	"shared/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type SsoService struct {
	VerifyEndpoint    string
	discoveryEndpoint string
	tenantId          string
	audienceId        string
	RedisRepo         *redis_repo.RedisRepo
}

type JWKS struct {
	Kid string
	X5c string
}

// func ctxListener(ctx context.Context) {
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				logrus.Error("======= DONE ")
// 				return
// 			default:
// 				logrus.Error("ctx continues ")
// 				time.Sleep(time.Second * 2)
// 			}
// 		}
// 	}()
// }

func NewSsoService() *SsoService {
	conf := config.Getenv()
	tenantId := conf.Azure.TenantId
	audienceId := conf.Azure.TenantId
	ctx := context.Background()
	// ctxListener(ctx)
	// * find the same value between :
	// decoded access token/header/kid AND
	// public keys / kid
	redisRepo := redis_repo.NewRedis(ctx)
	logrus.Warn(*redisRepo)
	return &SsoService{
		tenantId:          tenantId,
		audienceId:        audienceId,
		VerifyEndpoint:    "https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration",
		discoveryEndpoint: fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/keys?appid=%s", tenantId, audienceId),
		RedisRepo:         redisRepo,
	}
}

func (sso *SsoService) ExtractTokenClaims(tokenString string) (jwt.Claims, error) {
	keys, err := sso.getPublicKeys()
	if err != nil {
		return jwt.MapClaims{}, err
	}
	claims, ok := sso.isKidInPublicKeys(tokenString, keys)
	if !ok {
		return jwt.MapClaims{}, errors.New("token not valid")
	}
	return claims, nil
}

func (ssoService *SsoService) getRawAzureOpenId() ([]byte, error) {
	keysStr, err := ssoService.RedisRepo.Get("azure_Kids")
	if err != nil {
		logrus.Info("======> NO CACHE !!!")
		logrus.Warn("error message : ", err.Error())
	}
	if err == nil {
		logrus.Warn("=====> CACHE", keysStr)
		return []byte(keysStr), nil
	}

	resp, err := http.Get(ssoService.discoveryEndpoint)
	if err != nil {
		logrus.Info("error get redis : ", err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	ssoService.RedisRepo.Set("azure_Kids", string(body))
	return body, nil
}

func (ssoService *SsoService) getPublicKeys() ([]JWKS, error) {
	rawOpenIdKeys, err := ssoService.getRawAzureOpenId()
	// logrus.Info(rawOpenIdKeys)
	if err != nil {

		return []JWKS{}, err
	}

	microsoftDiscoveryKeyDto, err := utils.UnmarshallDto[dto_response.MicrosoftDiscoveryKeyDto](rawOpenIdKeys)
	if err != nil {
		return []JWKS{}, err
	}

	jwksKeys := make([]JWKS, 0, len(microsoftDiscoveryKeyDto.Keys))
	for _, key := range microsoftDiscoveryKeyDto.Keys {
		if key.Use == "sig" && key.Kty == "RSA" && len(key.Kid) > 0 && (len(key.X5c) > 0 || (len(key.N) > 0 && len(key.E) > 0)) {
			jwks := JWKS{
				Kid: key.Kid,
				X5c: key.X5c[0],
			}
			jwksKeys = append(jwksKeys, jwks)
		}
	}

	if len(jwksKeys) == 0 {
		return []JWKS{}, errors.New("no key found")
	}

	return jwksKeys, err
}

func (ssoService *SsoService) isKidInPublicKeys(tokenString string, keys []JWKS) (jwt.Claims, bool) {
	for _, key := range keys {
		token, _ := helper_jwt.Verify(tokenString)
		if token.Header["kid"] == key.Kid {
			return token.Claims, true
		}
	}
	return jwt.MapClaims{}, false
}
