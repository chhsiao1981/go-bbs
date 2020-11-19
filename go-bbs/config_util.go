package main

import (
	"strings"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(filename string) error {
	err := types.InitViper(filename)
	if err != nil {
		return err
	}
	log.Infof("viper keys: %v", viper.AllKeys())

	initConfig()

	ptttype.InitConfig()

	return nil
}

func setStringConfig(idx string, orig string) string {
	idx = "http-server." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}
