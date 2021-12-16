package openid

type Realm struct {
	Realm         string        `json:"realm"`
	Wellknown     Wellknown     `json:"wellknown"`
	Authorization Authorization `json:"authorization"`
	Token         Token         `json:"token"`
}
