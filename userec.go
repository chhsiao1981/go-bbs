package bbs

import (
	"github.com/PichuChen/go-bbs/ptttype"
	log "github.com/sirupsen/logrus"
)

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
type Userec struct {
	Version  uint32
	Userid   string
	Realname string
	Nickname string
	Passwd   string
	Pad1     uint8

	Uflag        uint32
	_unused1     uint32
	Userlevel    uint32
	Numlogindays uint32
	Numposts     uint32
	Firstlogin   uint32
	Lastlogin    uint32
	Lasthost     string
	// TODO
}

func NewUserecFromRaw(userecRaw *ptttype.UserecRaw) *Userec {
	log.Infof("userecRaw: %v", userecRaw)
	user := &Userec{}
	user.Version = userecRaw.Version
	user.Userid = CstrToString(userecRaw.UserID[:])
	user.Realname = Big5ToUtf8(CstrToBytes(userecRaw.RealName[:]))
	user.Nickname = Big5ToUtf8(CstrToBytes(userecRaw.Nickname[:]))
	user.Passwd = CstrToString(userecRaw.PasswdHash[:])
	user.Pad1 = userecRaw.Pad1

	user.Uflag = userecRaw.UFlag
	user._unused1 = userecRaw.Unused1
	user.Userlevel = userecRaw.UserLevel
	user.Numlogindays = userecRaw.NumLoginDays
	user.Numposts = userecRaw.NumPosts
	user.Firstlogin = uint32(userecRaw.FirstLogin)
	user.Lastlogin = uint32(userecRaw.LastLogin)
	user.Lasthost = CstrToString(userecRaw.LastHost[:])

	return user
}
