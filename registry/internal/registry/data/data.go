package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

func NewDB(ctx context.Context, conn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}

	db := &DB{
		Client: client,
	}

	db.init()

	return db, nil
}

func (db *DB) init() {
	db.Client.Database("registry")
}

func (d *DB) Close(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}
