package identity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"time"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

// Identity is the root for identity logic, no HTTP routing here.
type Identity struct {
	db         collections
	oAuth      *oauth2.Config
	privateKey *ecdsa.PrivateKey
}

type collections struct {
	users *mongo.Collection
}

func New(privateKey *ecdsa.PrivateKey, oAuth config.OAuthConfig, db *data.DB) *Identity {
	return &Identity{
		db: collections{
			users: db.Client.Database("registry").Collection("users"),
		},
		oAuth: &oauth2.Config{
			ClientID:     oAuth.ClientID,
			ClientSecret: oAuth.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oAuth.AuthorizationURL.String(),
				TokenURL: oAuth.TokenURL.String(),
			},
			RedirectURL: oAuth.RedirectURL.String(),
			Scopes:      []string{"identify", "email"},
		},
		privateKey: privateKey,
	}
}

// Expose OAuth config for HTTP layer
func (s *Identity) OAuthConfig() *oauth2.Config {
	return s.oAuth
}

func (s *Identity) IdentityCallback(ctx context.Context, id string) (string, error) {
	user, err := s.newUser(ctx, id)
	if err != nil {
		return "", err
	}

	tok, err := s.GenerateJWT(user)
	if err != nil {
		fmt.Printf("failed to generate jwt: %v\n", err)
		return "", err
	}

	return tok, nil
}

func (s *Identity) newUser(ctx context.Context, id string) (*User, error) {
	db := s.db.users
	filter := bson.M{"discord_id": id}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "lastLogin", Value: time.Now()}}}}
	result := db.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true))

	user := new(User)
	if err := result.Decode(&user); err != nil {
		slog.Error("decoding find one and update result", "err", err.Error())
		return nil, err
	}

	slog.Info("new user upserted", "id", user.ID)
	return user, nil
}

func (s *Identity) GenerateJWT(user *User) (string, error) {
	b, err := user.ID.MarshalText()
	if err != nil {
		return "", err
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": string(b),
	})

	return tok.SignedString(s.privateKey)
}

// PrivateKey returns the ECDSA private key for this identity.
func (s *Identity) PrivateKey() *ecdsa.PrivateKey {
	return s.privateKey
}
