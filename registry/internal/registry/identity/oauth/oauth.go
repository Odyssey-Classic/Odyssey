package oauth

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthServer struct {
	Config           *oauth2.Config
	IdentityCallback func(context.Context, string) (string, error)
}

func (s *OAuthServer) Run() {
	http.HandleFunc("/oauth/callback", s.oAuthCallback)

	srv := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func (s *OAuthServer) OAuthRedirect(w http.ResponseWriter, r *http.Request) {
	// TODO: save verifier to use it in callback
	// verifier := oauth2.GenerateVerifier()
	url := s.Config.AuthCodeURL("state") //, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusFound)
}
