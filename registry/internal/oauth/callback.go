package oauth

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"

	"golang.org/x/oauth2"
)

func RunCallback(cfg *oauth2.Config, verifier string) {
	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "no code", http.StatusBadRequest)
			return
		}

		// var code string
		// if _, err := fmt.Scan(&code); err != nil {
		// 	log.Fatal(err)
		// }
		// slog.Info("", "code", code)
		// tok, err := cfg.Exchange(ctx, code, oauth2.VerifierOption(verifier))
		// if err != nil {
		// 	log.Fatal(err)
		// }

		w.Write([]byte("ok"))
		tok, err := cfg.Exchange(ctx, code, oauth2.VerifierOption(verifier))
		if err != nil {
			slog.Error("failed exchange", "error", err)
			return
		}

		slog.Info("", "code", code)
		client := cfg.Client(ctx, tok)
		res, err := client.Get("https://discord.com/api/users/@me")
		if err != nil {
			log.Fatal(err)
		}
		msg, _ := io.ReadAll(res.Body)
		fmt.Println(string(msg))
	})

	s := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(s.ListenAndServe())
}
