package ports

import "github.com/pedrolopesme/keyghost/internal/core/domain"

type OpenIdGhostContract interface {
	Wellknown(profile domain.ServerProfile) (string, error)
}
