package servers

import (
	"context"
	"testing"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServersTestSuite struct {
	suite.Suite
	MongoClient *mongo.Client
	Service     *Service
	Resource    *dockertest.Resource
	Pool        *dockertest.Pool
}

func (suite *ServersTestSuite) SetupSuite() {
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
	dsn := "mongodb://localhost:" + resource.GetPort("27017/tcp")
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
	db := &data.DB{Client: client}
	suite.Service = NewService(db)
}

func (suite *ServersTestSuite) TearDownSuite() {
	if suite.MongoClient != nil {
		_ = suite.MongoClient.Disconnect(context.Background())
	}
	if suite.Pool != nil && suite.Resource != nil {
		_ = suite.Pool.Purge(suite.Resource)
	}
}

func (suite *ServersTestSuite) TestFindByID() {
	ctx := context.Background()
	// Insert a server directly into the servers collection
	coll := suite.MongoClient.Database("registry").Collection("servers")
	insertResult, err := coll.InsertOne(ctx, map[string]interface{}{
		"name": "TestServer",
		"key":  "test-key-hash",
		"user": "test-user-id",
	})
	suite.Require().NoError(err)
	suite.NotNil(insertResult.InsertedID)

	id := insertResult.InsertedID
	idStr := ""
	if oid, ok := id.(interface{ Hex() string }); ok {
		idStr = oid.Hex()
	} else {
		idStr = id.(string)
	}

	// Find by ID using the Service
	servers, err := suite.Service.FindByID(ctx, idStr)
	suite.Require().NoError(err)
	suite.Len(servers, 1)
	suite.Equal("TestServer", servers[0].Name)
}

func (suite *ServersTestSuite) TestResetAPIKey_OnlyAffectsCorrectServer() {
	service := suite.Service
	ctx := context.Background()

	user1 := &identity.User{ID: primitive.NewObjectID()}
	user2 := &identity.User{ID: primitive.NewObjectID()}

	// Register two servers
	id1, key1, err := service.RegisterServer(ctx, "Server1", user1)
	suite.Require().NoError(err)
	id2, _, err := service.RegisterServer(ctx, "Server2", user2)
	suite.Require().NoError(err)

	// Get original hashes
	servers1, err := service.FindByID(ctx, id1)
	suite.Require().NoError(err)
	suite.Require().Len(servers1, 1)
	origHash1 := servers1[0].Key

	servers2, err := service.FindByID(ctx, id2)
	suite.Require().NoError(err)
	suite.Require().Len(servers2, 1)
	origHash2 := servers2[0].Key

	// Reset API key for server 1
	newKey, err := service.ResetAPIKey(ctx, id1)
	suite.Require().NoError(err)
	suite.Require().NotEqual(key1, newKey)

	// Check server 1's key changed
	servers1After, err := service.FindByID(ctx, id1)
	suite.Require().NoError(err)
	suite.Require().Len(servers1After, 1)
	suite.Require().NotEqual(origHash1, servers1After[0].Key)

	// Check server 2's key did not change
	servers2After, err := service.FindByID(ctx, id2)
	suite.Require().NoError(err)
	suite.Require().Len(servers2After, 1)
	suite.Require().Equal(origHash2, servers2After[0].Key)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServersTestSuite))
}
