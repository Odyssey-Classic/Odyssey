package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string `bson:"_id,omitempty"`
	DiscordID string `bson:"discord_id"`
}

func main() {
	// Uses the same environment variable as the registry for MongoDB URI: ODY_REG_DB_CONNECTION
	uri := os.Getenv("ODY_REG_DB_CONNECTION")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	coll := client.Database("registry").Collection("users")
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}
	defer cur.Close(ctx)

	fmt.Println("Users:")
	for cur.Next(ctx) {
		var user User
		if err := cur.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v", err)
			continue
		}
		fmt.Printf("ID: %v, DiscordID: %v\n", user.ID, user.DiscordID)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
}
