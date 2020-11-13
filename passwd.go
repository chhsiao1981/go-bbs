package bbs

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"github.com/PichuChen/go-bbs/ptt"
	"github.com/PichuChen/go-bbs/ptttype"
)

func Login(userID string, passwd string, ip string) (*Userec, error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))
	passwdRaw := []byte(passwd)
	ipRaw := &[ptttype.IPV4LEN + 1]byte{}
	copy(ipRaw[:], []byte(ip))

	userRaw, err := ptt.LoginQuery(userIDRaw, passwdRaw, ipRaw)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromRaw(userRaw)

	return user, nil
}

func OpenUserecFile(filename string) ([]*Userec, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*Userec{}

	for {
		user, eachErr := NewUserecWithFile(file)
		if eachErr != nil {
			// io.EOF is reading correctly to the end the file.
			if eachErr == io.EOF {
				break
			}

			err = eachErr
			break
		}
		ret = append(ret, user)
	}

	return ret, err

}

func NewUserecWithFile(file *os.File) (*Userec, error) {
	userecRaw := &ptttype.UserecRaw{}

	err := binary.Read(file, binary.LittleEndian, userecRaw)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromRaw(userecRaw)

	return user, nil
}
