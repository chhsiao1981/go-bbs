package types

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func InitConfig() (err error) {
	config()
	return postConfig()
}

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

//setTimeLocation
//
//
func setTimeLocation(timeLocation string) (origTimeLocation string, err error) {
	origTimeLocation = TIME_LOCATION
	TIME_LOCATION = timeLocation

	TIMEZONE, err = time.LoadLocation(TIME_LOCATION)

	return origTimeLocation, err
}

func postConfig() (err error) {
	_, err = setTimeLocation(TIME_LOCATION)
	if err != nil {
		return err
	}

	return nil
}
