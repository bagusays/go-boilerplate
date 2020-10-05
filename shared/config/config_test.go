package config_test

import (
	"go-boilerplate/shared/config"
	"os"
	"testing"
)

func init() {
	os.Chdir("../../")
}

func TestConfig(t *testing.T) {
	config.LoadConfig()
	conf := config.GetConfig()
	if conf.GetJWTSecretKey() == "" {
		t.Error("JWT is empty")
	}
}
