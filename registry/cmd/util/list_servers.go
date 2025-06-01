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

type Server struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

func main() {
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

	coll := client.Database("registry").Collection("servers")
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to query servers: %v", err)
	}
	defer cur.Close(ctx)

	fmt.Println("Servers:")
	for cur.Next(ctx) {
		var server Server
		if err := cur.Decode(&server); err != nil {
			log.Printf("Failed to decode server: %v", err)
			continue
		}
		fmt.Printf("ID: %v, Name: %v\n", server.ID, server.Name)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
}
