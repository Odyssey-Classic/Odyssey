package oauth

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"strings"

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

		verifiers: make(map[string]string, 10),
	}
}

func (s *OAuthServer) OAuthRedirect(w http.ResponseWriter, r *http.Request) {
	state, verifier := s.createVerifier(r)
	url := s.Config.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusFound)
}

func (s *OAuthServer) createVerifier(r *http.Request) (string, string) {
	// TODO: How do we delete UNUSED verifiers? If a user doesn't complete the oauth flow.

	// hash just the ip address, without remote port
	ipHash := md5.Sum([]byte(strings.Split(r.RemoteAddr, ":")[0]))

	state := base64.StdEncoding.EncodeToString(ipHash[:])
	verifier := oauth2.GenerateVerifier()
	s.verifiers[string(state)] = verifier
	return state, verifier
}
