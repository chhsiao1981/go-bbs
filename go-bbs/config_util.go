package main

import (
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}
