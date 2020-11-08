package bbs

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func getRandom() uuid.UUID {
	return uuid.New()
}

func getCurrentMilliTS() int64 {
	return time.Now().UnixNano() / 1000000
}

func isRegularFile(filename string) bool {
	// dashf in pttbbs
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsRegular()
}

func readFile(filename string) ([]byte, error) {
	if !isRegularFile(filename) {
		return nil, errors.New("not regular")
	}
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return nil, err
	}

	return data, nil
}

func CopyFile(src string, dst string) (int64, error) {
	if !isRegularFile(src) {
		return 0, os.ErrInvalid
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func binRead(file *os.File, v interface{}, theSize uintptr) error {
	err := binary.Read(file, binary.LittleEndian, v)
	if err != nil {
		return err
	}
	binSize := binary.Size(v)
	nToRead := int64(theSize) - int64(binSize)
	if nToRead < 0 {
		log.WithFields(log.Fields{"theSize": theSize, "binSize": binSize}).Warn("binRead: theSize < binSize")
		return nil
	}
	log.WithFields(log.Fields{"theSize": theSize, "binSize": binSize, "nToRead": nToRead}).Debug("to seek")
	_, err = file.Seek(nToRead, 1)
	if err != nil {
		return err
	}
	return nil
}

func binWrite(file *os.File, v interface{}, theSize uintptr) error {
	err := binary.Write(file, binary.LittleEndian, v)
	if err != nil {
		return err
	}
	binSize := binary.Size(v)
	nToWrite := int64(theSize) - int64(binSize)
	if nToWrite < 0 {
		log.WithFields(log.Fields{"theSize": theSize, "binSize": binSize}).Warn("binWrite: theSize < binSize")
		return nil
	}
	log.WithFields(log.Fields{"theSize": theSize, "binSize": binSize, "nToWrite": nToWrite}).Debug("to write dummy")
	dummy := make([]byte, nToWrite)
	_, err = file.Write(dummy)
	if err != nil {
		return err
	}
	return nil
}
