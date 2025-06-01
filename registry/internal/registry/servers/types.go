package servers

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServerInfo holds metadata about a registered game server.
type ServerInfo struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Key  string             `bson:"key" json:"key"`
	Name string             `bson:"name" json:"name"`
	User primitive.ObjectID `bson:"user" json:"user"`
}
