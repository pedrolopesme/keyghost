package ports

type OpenIdGhostContract interface {
	Wellknown(profile domain.ServerProfile) (string, error)
}
