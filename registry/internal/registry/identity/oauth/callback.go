package oauth

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"golang.org/x/oauth2"
)

type DiscordIdentity struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (s *OAuthServer) OAuthCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cfg := s.Config

	state := r.URL.Query().Get("state")
	if state == "" {
		http.Error(w, "no state", http.StatusBadRequest)
		return
	}

	verifier := s.verifiers[state]
	if verifier == "" {
		http.Error(w, "no verifier", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "no code", http.StatusBadRequest)
		return
	}

	tok, err := cfg.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		slog.Error("failed exchange", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info("", "code", code)
	client := cfg.Client(ctx, tok)
	res, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		slog.Error("failed to get user", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, _ := io.ReadAll(res.Body)

	var user DiscordIdentity
	err = json.Unmarshal(msg, &user)
	if err != nil {
		slog.Error("failed to unmarshal", "error", err)
	}

	jwtToken, err := s.IdentityCallback(r.Context(), user.Id)
	if err != nil {
		slog.Error("failed to get jwt", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jwtToken))
}
