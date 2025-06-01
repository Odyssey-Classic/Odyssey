package identity

import (
	"context"
	"fmt"
	"testing"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IdentityTestSuite struct {
	suite.Suite
	MongoClient *mongo.Client
	Identity    *Identity
	Resource    *dockertest.Resource
	Pool        *dockertest.Pool
}

func (suite *IdentityTestSuite) SetupSuite() {
	pool, err := dockertest.NewPool("")
	suite.Require().NoError(err)
	suite.Pool = pool

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "7",
		Env:        []string{"MONGO_INITDB_DATABASE=registry"},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	suite.Require().NoError(err)
	suite.Resource = resource

	var client *mongo.Client
	dsn := fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp"))
	err = pool.Retry(func() error {
		var err error
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
		if err != nil {
			return err
		}
		return client.Ping(context.Background(), nil)
	})
	suite.Require().NoError(err)
	suite.MongoClient = client
	suite.Identity = &Identity{db: collections{users: data.DB{Client: client}.Client.Database("registry").Collection("users")}}
}

func (suite *IdentityTestSuite) TearDownSuite() {
	if suite.MongoClient != nil {
		_ = suite.MongoClient.Disconnect(context.Background())
	}
	if suite.Pool != nil && suite.Resource != nil {
		_ = suite.Pool.Purge(suite.Resource)
	}
}

// func (suite *IdentityTestSuite) TestFindUserByDiscordId() {
// 	coll := suite.MongoClient.Database("registry").Collection("users")
// 	_, err := coll.InsertOne(context.Background(), map[string]interface{}{"discord_id": "test123"})
// 	suite.Require().NoError(err)
// 	user, err := suite.Identity.FindUserByDiscordId(context.Background(), "test123")
// 	suite.Require().NoError(err)
// 	suite.Equal("test123", user.DiscordID)
// }

func TestIdentityTestSuite(t *testing.T) {
	suite.Run(t, new(IdentityTestSuite))
}
