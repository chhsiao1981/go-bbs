package bbs

import (
	"github.com/PichuChen/go-bbs/ptt"
	"github.com/PichuChen/go-bbs/ptttype"
)

func Register(
	userID string,
	passwd string,
	ip string,
	email string,

	nickname string,
	realname string,
	career string,
	address string,
	over18 bool,
) (user *Userec, err error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))

	passwdRaw := []byte(passwd)

	ipRaw := &[ptttype.IPV4LEN + 1]byte{}
	copy(ipRaw[:], []byte(ip))

	emailRaw := &[ptttype.EMAILSZ]byte{}
	copy(emailRaw[:], []byte(email))

	nicknameRaw := &[ptttype.NICKNAMESZ]byte{}
	copy(nicknameRaw[:], []byte(nickname))

	realnameRaw := &[ptttype.REALNAMESZ]byte{}
	copy(realnameRaw[:], []byte(realname))

	careerRaw := &[ptttype.CAREERSZ]byte{}
	copy(careerRaw[:], []byte(career))

	addressRaw := &[ptttype.ADDRESSSZ]byte{}
	copy(addressRaw[:], []byte(address))

	userRaw, err := ptt.NewRegister(
		userIDRaw,
		passwdRaw,
		ipRaw,
		emailRaw,
		false,
		false,

		nicknameRaw,
		realnameRaw,
		careerRaw,
		addressRaw,
		over18,
	)
	if err != nil {
		return nil, err
	}

	user = NewUserecFromRaw(userRaw)

	return user, nil
}
