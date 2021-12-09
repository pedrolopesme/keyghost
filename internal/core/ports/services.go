package ports

import "github.com/pedrolopesme/keyghost/internal/core/domain"

type OpenIdGhostPort interface {
	Wellknown(profile domain.ServerProfile) (string, error)
}
type ServerProfilePort interface {
	LoadServerProfile(path string) (*domain.ServerProfile, error)
}
