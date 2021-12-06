package openid

type Wellknown struct {
	Issuer                 string   `json:"open"`
	AuthorizationEndpoint  string   `json:"authorization_endpoint"`
	TokenEndpoint          string   `json:"token_endpoint"`
	UserInfoEndpoint       string   `json:"userinfo_endpoint"`
	EndSessionEndpoint     string   `json:"end_session_endpoint"`
	JwksUri                string   `json:"jwks_uri"`
	GrantTypesSupported    []string `json:"grant_types_supported"`
	ResponseTypesSupported []string `json:"response_types_supported"`
	SubjectTypesSupported  []string `json:"subject_types_supported"`
	IdTokenSignign         []string `json:"id_token_signing_alg_values_supported"`
	ResponseModesSupported []string `json:"response_modes_supported"`
}
