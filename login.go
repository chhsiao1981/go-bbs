package bbs

import (
	"github.com/PichuChen/go-bbs/ptt"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/sirupsen/logrus"
)

//Login
//
//XXX need to check for the permission
func Login(userID string, passwd string, ip string) (*Userec, error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))
	passwdRaw := []byte(passwd)
	ipRaw := &[ptttype.IPV4LEN + 1]byte{}
	copy(ipRaw[:], []byte(ip))

	userRaw, err := ptt.LoginQuery(userIDRaw, passwdRaw, ipRaw)
	logrus.Debugf("bbs.passwd.Login: after LoginQuery: userRaw: %v e: %v", userRaw, err)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromRaw(userRaw)

	return user, nil
}
