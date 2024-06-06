package oauth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthServer struct {
	Config           *oauth2.Config
	IdentityCallback func(context.Context, string) (string, error)

	verifiers map[string]string
}

func New(config *oauth2.Config, callback func(context.Context, string) (string, error)) *OAuthServer {
	return &OAuthServer{
		Config:           config,
		IdentityCallback: callback,
		verifiers:        make(map[string]string, 10),
	}
}

func (s *OAuthServer) OAuthRedirect(w http.ResponseWriter, r *http.Request) {
	// TODO: save verifier to use it in callback
	state, verifier := s.createVerifier()
	url := s.Config.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusFound)
}

func (s *OAuthServer) createVerifier() (string, string) {
	state := oauth2.GenerateVerifier()
	verifier := oauth2.GenerateVerifier()
	s.verifiers[state] = verifier
	return state, verifier
}
