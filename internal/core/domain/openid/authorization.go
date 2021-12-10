package openid

type Authorization struct {
	RedirectTo string `json:"redirect_to"`
	Code       string `json:"code"`
}
