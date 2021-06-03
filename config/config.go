package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")

	viper.SetConfigName("config")

	if env := viper.Get("ENV"); env == "local" {
		viper.SetConfigName("config-local")
		viper.AddConfigPath("./..")
	}
}

func Init() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found at specified paths")
		} else {
			fmt.Printf("fatal error config file: %s ", err)
		}
	}
}
