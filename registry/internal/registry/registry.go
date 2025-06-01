package registry

import (
	"crypto/ecdsa"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
	"golang.org/x/oauth2"
)

type Registry struct {
	db       *data.DB
	identity *identity.Identity
	servers  *servers.Service
}

// OAuthConfig returns the registry's OAuth config via the IdentityServer.
func (r *Registry) OAuthConfig() *oauth2.Config {
	return r.identity.OAuthConfig()
}

// PrivateKey should not be modified after creation.
func (r *Registry) PrivateKey() *ecdsa.PrivateKey {
	return r.identity.PrivateKey()
}

func NewRegistry(db *data.DB, oauthConfig config.OAuthConfig, privateKey *ecdsa.PrivateKey) *Registry {
	return &Registry{
		db:       db,
		identity: identity.New(privateKey, oauthConfig, db),
		servers:  servers.NewService(db),
	}
}

func (r *Registry) ServersService() *servers.Service {
	// For now, create a new Service each time. You may want to store this in Registry if you want persistence.
	return r.servers
}

func (r *Registry) IdentityService() *identity.Identity {
	return r.identity
}
