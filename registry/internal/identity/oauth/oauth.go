package oauth

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthServer struct {
	Config           *oauth2.Config
	IdentityCallback func(string) (string, error)
}

func (s *OAuthServer) Run() {
	// verifier := oauth2.GenerateVerifier()

	// TODO: save verifier to use it in callback
	url := s.Config.AuthCodeURL("state") //, oauth2.S256ChallengeOption(verifier))

	http.HandleFunc("/oauth/callback", s.oAuthCallback)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	})

	srv := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}
