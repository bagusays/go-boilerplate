package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	imOnce sync.Once
)

func LoadConfig() {
	imOnce.Do(func() {
		viper.SetConfigName("app.config.yaml") // name of config file (without extension)
		viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")               // optionally look for config in the working directory
		err := viper.ReadInConfig()            // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
}
