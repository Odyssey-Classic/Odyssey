package oauth

import (
	"fmt"

	"golang.org/x/oauth2"
)

func Run(cfg *oauth2.Config) {
	// ctx := context.Background()

	verifier := oauth2.GenerateVerifier()

	url := cfg.AuthCodeURL("state", oauth2.S256ChallengeOption(verifier))
	fmt.Printf("auth: %v\n", url)

	go RunCallback(cfg, verifier)

	for {
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

	// client := cfg.Client(ctx, tok)
	// res, err := client.Get("https://discord.com/api/users/@me")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(res.Body)
}
