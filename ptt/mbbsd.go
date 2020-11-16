package ptt

import (
	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/names"
	"github.com/PichuChen/go-bbs/ptttype"
	log "github.com/sirupsen/logrus"
)

//LoginQuery
//Params
//	userID: userID
//	passwd: passwd
//	ip: ip
//
//Return
//	*UserecRaw: user
//  error: err
func LoginQuery(userID *[ptttype.IDLEN + 1]byte, passwd []byte, ip *[ptttype.IPV4LEN + 1]byte) (*ptttype.UserecRaw, error) {
	if !names.IsValidUserID(userID) {
		return nil, ptttype.ErrInvalidUserID
	}

	_, cuser, err := initCurrentUser(userID)
	log.Infof("after initCurrentUser: cuser: %v e: %v", cuser, err)
	if err != nil {
		return nil, err
	}

	isValid, err := cmbbs.CheckPasswd(cuser.PasswdHash[:], passwd)
	log.Infof("mbbsd.LoginQuery: after CheckPasswd: isValid: %v e: %v", isValid, err)
	if err != nil {
		cmbbs.LogAttempt(userID, ip, true)
		return nil, err
	}

	if !isValid {
		cmbbs.LogAttempt(userID, ip, true)
		return nil, ptttype.ErrInvalidUserID
	}

	return cuser, nil
}
