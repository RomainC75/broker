package dto_response

type MicrosoftDiscoveryKeyDto struct {
	Keys []KeyDto `json:"keys"`
}

type KeyDto struct {
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
	N   string   `json:"n"`
	E   string   `json:"E"`
	X5c []string `json:"x5c"`
}
