package config_test

import (
	"go-boilerplate/shared/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	res := config.New("../../configs/")

	assert.Greater(t, len(res.Apps.Name), 1)
	assert.NotEqual(t, "", res.Apps.Port)
}
