package types

import "time"

func InitConfig() (err error) {
	TIME_LOCATION = setStringConfig("TIME_LOCATION", TIME_LOCATION)

	return postConfig()
}

func setTimeLocation() (err error) {
	TIMEZONE, err = time.LoadLocation(TIME_LOCATION)
	if err != nil {
		return err
	}

	return nil
}

func postConfig() (err error) {
	err = setTimeLocation()
	if err != nil {
		return err
	}

	return nil
}
