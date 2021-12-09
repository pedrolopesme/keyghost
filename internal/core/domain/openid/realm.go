package openid

type Realm struct {
	Realm     string    `json:"realm"`
	Wellknown Wellknown `json:"wellknown"`
}
