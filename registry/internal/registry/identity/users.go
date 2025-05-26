package identity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DiscordID string             `bson:"discord_id"`
}

func (s *Identity) FindUserByDiscordId(ctx context.Context, discordID string) (*User, error) {
	user := &User{
		DiscordID: discordID,
	}
	res := s.db.Client.Database("registry").Collection("users").FindOne(ctx, user)
}
