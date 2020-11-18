package types

import (
	"strings"

	"github.com/spf13/viper"
)

func setStringConfig(idx string, orig string) string {
	idx = "types." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}

func setBoolConfig(idx string, orig bool) bool {
	idx = "types." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetBool(idx)
}

func setIntConfig(idx string, orig int) int {
	idx = "types." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}
	return viper.GetInt(idx)
}

func setDoubleConfig(idx string, orig float64) float64 {
	idx = "types." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetFloat64(idx)
}
