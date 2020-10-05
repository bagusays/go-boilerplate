package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type (
	ImmutableConfigInterface interface {
		GetJWTSecretKey() string
		GetServiceName() string
	}

	im struct {
		JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
		ServiceName  string `mapstructure:"SERVICE_NAME"`
	}
)

func (i *im) GetJWTSecretKey() string {
	return i.JWTSecretKey
}

func (i *im) GetServiceName() string {
	return i.ServiceName
}

var (
	imOnce sync.Once
	config im
)

func LoadConfig() {
	imOnce.Do(func() {
		v := viper.New()

		if flag.Lookup("test.v") != nil {
			v.SetConfigName("app.config.test")
			fmt.Println("test")
		} else {
			appEnv, exists := os.LookupEnv("APP_ENV")
			if exists {
				if appEnv == "staging" {
					v.SetConfigName("app.config.staging")
					fmt.Println("staging")
				} else if appEnv == "production" {
					v.SetConfigName("app.config.prod")
					fmt.Println("prod")
				}
			} else {
				v.SetConfigName("app.config.dev")
				fmt.Println("dev")
			}
		}

		v.AddConfigPath(".") // optionally look for config in the working directory
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("[CONFIG][missing] Failed to read app.config.* file. %v", err)
		}

		err := v.Unmarshal(&config)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	})
}

func GetConfig() ImmutableConfigInterface {
	return &config
}
