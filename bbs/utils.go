package bbs

import (
	"errors"
	"io/ioutil"
	"os"
)

func dashf(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsRegular()
}

func readFile(filename string) ([]byte, error) {
	if !dashf(filename) {
		return nil, errors.New("not regular")
	}
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return nil, err
	}

	return data, nil
}
