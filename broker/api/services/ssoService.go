package services

import (
	dto_response "broker/api/dto/response"
	helper_jwt "broker/helpers/jwt"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shared/config"
	"shared/utils"
)

type SsoService struct {
	VerifyEndpoint    string
	discoveryEndpoint string
	tenantId          string
	audienceId        string
}

type JWKS struct {
	Kid string
	X5c string
}

func NewSsoService() *SsoService {
	conf := config.Getenv()
	tenantId := conf.Azure.TenantId
	audienceId := conf.Azure.TenantId

	// * find the same value between :
	// decoded access token/header/kid AND
	// public keys / kid

	return &SsoService{
		tenantId:          tenantId,
		audienceId:        audienceId,
		VerifyEndpoint:    "https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration",
		discoveryEndpoint: fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/keys?appid=%s", tenantId, audienceId),
	}

}

func (ssoService *SsoService) GetPublicKeys(token string) ([]JWKS, error) {
	resp, err := http.Get(ssoService.discoveryEndpoint)
	if err != nil {
		return []JWKS{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	microsoftDiscoveryKeyDto, err := utils.UnmarshallDto[dto_response.MicrosoftDiscoveryKeyDto](body)
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
	utils.PrettyDisplay(" RECEIVED KID KEYS : ", jwksKeys)
	ssoService.IsKidInPublicKeys(token, jwksKeys)

	return jwksKeys, err
}

func (ssoService *SsoService) IsKidInPublicKeys(token string, keys []JWKS) {
	for range keys {

		helper_jwt.Verify(token)
	}
}
