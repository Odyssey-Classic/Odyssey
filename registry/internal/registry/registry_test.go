package registry_test

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"crypto/elliptic"
// 	"crypto/rand"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/suite"

// 	"github.com/FosteredGames/Odyssey/registry/internal/config"
// 	"github.com/FosteredGames/Odyssey/registry/internal/registry"
// )

// type RegistryTestSuite struct {
// 	suite.Suite
// 	reg *registry.Registry
// }

// func (s *RegistryTestSuite) SetupTest() {
// 	s.reg = &registry.Registry{
// 		DB:          nil, // You can mock or stub DB if needed
// 		OAuthConfig: config.OAuthConfig{},
// 		PrivateKey:  generateTestKey(),
// 	}
// }

// func generateTestKey() *ecdsa.PrivateKey {
// 	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	return key
// }

// func (s *RegistryTestSuite) TestRegistryGracefulShutdown() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	done := make(chan error, 1)
// 	go func() {
// 		err := s.reg.Run(ctx)
// 		done <- err
// 	}()

// 	time.Sleep(100 * time.Millisecond) // Give the server time to start
// 	cancel()                           // Trigger shutdown

// 	select {
// 	case err := <-done:
// 		s.True(err == nil || err == http.ErrServerClosed, "expected nil or http.ErrServerClosed, got: %v", err)
// 	case <-time.After(2 * time.Second):
// 		s.Fail("Registry did not shut down in time")
// 	}
// }

// func TestRegistryTestSuite(t *testing.T) {
// 	suite.Run(t, new(RegistryTestSuite))
// }
