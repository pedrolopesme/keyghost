package domain

import "github.com/pedrolopesme/keyghost/internal/core/domain/openid"

type ServerProfile struct {
	ProfileName string           `json:"profile_name"`
	Realm       string           `json:"realm"`
	Wellknown   openid.Wellknown `json:"wellknown"`
}
