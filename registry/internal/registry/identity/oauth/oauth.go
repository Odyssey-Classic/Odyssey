package oauth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthServer struct {
	Config           *oauth2.Config
	IdentityCallback func(context.Context, string) (string, error)
}

func (s *OAuthServer) OAuthRedirect(w http.ResponseWriter, r *http.Request) {
	// TODO: save verifier to use it in callback
	// verifier := oauth2.GenerateVerifier()
	url := s.Config.AuthCodeURL("state") //, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusFound)
}
