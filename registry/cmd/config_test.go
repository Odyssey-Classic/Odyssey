package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFromEnv(t *testing.T) {
	t.Run("missing ODY_AUTHORIZATION_URL", func(t *testing.T) {
		c, err := ConfigFromEnv()
		assert.Nil(t, c)
		assert.NotNil(t, err)
		t.Log(err)
	})
}
