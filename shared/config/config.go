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
		GetServiceName() string
		GetJWTSecretKey() string
		GetDatabaseMySQLHost() string
		GetDatabaseMySQLUsername() string
		GetDatabaseMySQLPassword() string
		GetDatabaseMySQLDBName() string
	}

	im struct {
		ServiceName           string `mapstructure:"SERVICE_NAME"`
		JWTSecretKey          string `mapstructure:"JWT_SECRET_KEY"`
		DatabaseMySQLHost     string `mapstructure:"DATABASE_MYSQL_HOST"`
		DatabaseMySQLUsername string `mapstructure:"DATABASE_MYSQL_USERNAME"`
		DatabaseMySQLPassword string `mapstructure:"DATABASE_MYSQL_PASSWORD"`
		DatabaseMySQLDBName   string `mapstructure:"DATABASE_MYSQL_DB_NAME"`
	}
)

func (i *im) GetServiceName() string {
	return i.ServiceName
}

func (i *im) GetJWTSecretKey() string {
	return i.JWTSecretKey
}

func (i *im) GetDatabaseMySQLHost() string {
	return i.DatabaseMySQLHost
}

func (i *im) GetDatabaseMySQLUsername() string {
	return i.DatabaseMySQLUsername
}

func (i *im) GetDatabaseMySQLPassword() string {
	return i.DatabaseMySQLPassword
}

func (i *im) GetDatabaseMySQLDBName() string {
	return i.DatabaseMySQLDBName
}

var (
	imOnce sync.Once
	config im
)

func LoadConfig() {
	imOnce.Do(func() {
		v := viper.New()
		var configEnv string

		if flag.Lookup("test.v") != nil {
			v.SetConfigName("app.config.test")
			configEnv = "test"
		} else {
			appEnv, exists := os.LookupEnv("APP_ENV")
			if exists {
				if appEnv == "staging" {
					v.SetConfigName("app.config.staging")
					configEnv = "staging"
				} else if appEnv == "production" {
					v.SetConfigName("app.config.prod")
					configEnv = "prod"
				}
			} else {
				v.SetConfigName("app.config.dev")
				configEnv = "dev"
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

		fmt.Println(fmt.Sprintf("[CONFIG ENVIRONMENT] %s", configEnv))
	})
}

func GetConfig() ImmutableConfigInterface {
	return &config
}
