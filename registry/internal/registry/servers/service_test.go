package servers

import (
	"context"
	"testing"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceTestSuite struct {
	suite.Suite
	MongoClient *mongo.Client
	Service     *Service
	Resource    *dockertest.Resource
	Pool        *dockertest.Pool
}

func (suite *ServiceTestSuite) SetupSuite() {
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

func (suite *ServiceTestSuite) TearDownSuite() {
	if suite.MongoClient != nil {
		_ = suite.MongoClient.Disconnect(context.Background())
	}
	if suite.Pool != nil && suite.Resource != nil {
		_ = suite.Pool.Purge(suite.Resource)
	}
}

func (suite *ServiceTestSuite) TestFindByID() {
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

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
