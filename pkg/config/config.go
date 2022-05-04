package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

func Load() {

	if runtime.GOOS == "linux" {
		viper.AddConfigPath("/usr/local/etc/url-sh")
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Could not read config file: ", err)
	}
}
