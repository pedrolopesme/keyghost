package domain

import "github.com/pedrolopesme/keyghost/internal/core/domain/openid"

type ServerProfile struct {
	ProfileName string         `json:"profile_name"`
	Realms      []openid.Realm `json:"realms"`
}

func (s ServerProfile) Realm(name string) *openid.Realm {
	for i := range s.Realms {
		if s.Realms[i].Realm == name {
			return &s.Realms[i]
		}
	}
	return nil
}
