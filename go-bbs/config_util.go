package main

import (
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {
	initConfig()
	return nil
}

func setStringConfig(idx string, orig string) string {
	idx = "http-server." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}
