package openid

type TokensResponse struct {
	AcessToken       string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	IdToken          string `json:"id_token"`
	NotBeforePolice  int    `json:"not_before_police"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"openid"`
}

type AuthorizationCode struct {
	Code     string         `json:"code"`
	Response TokensResponse `json:"response"`
}

type Token struct {
	AuthorizationCode AuthorizationCode `json:"authorization_code"`
}
