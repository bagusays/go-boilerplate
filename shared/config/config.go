package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Apps     Apps     `json:"apps"`
	Database Database `json:"database"`
	Log      Log      `json:"log"`
}

type Log struct {
	IsWriteToFile bool `json:"isWriteToFile"`
}

type Apps struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(configPath string) *Config {
	env := "dev"
	if v, ok := os.LookupEnv("ENV"); ok {
		env = v
	} else if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		env = "test"
	}

	path := fmt.Sprintf("%sconfig.%s.json", configPath, env)
	fmt.Println("Loading config", path)

	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}
