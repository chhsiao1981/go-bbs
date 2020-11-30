package ptt

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
)

func initCurrentUser(userID *ptttype.UserID_t) (uid int32, user *ptttype.UserecRaw, err error) {
	return cmbbs.PasswdLoadUser(userID)
}

func passwdSyncUpdate(uid int32, user *ptttype.UserecRaw) error {
	if uid < 1 || uid > ptttype.MAX_USERS {
		return cache.ErrInvalidUID
	}

	user.Money = cache.MoneyOf(uid)

	err := cmbbs.PasswdUpdate(uid, user)
	if err != nil {
		return err
	}

	return nil
}

func passwdSyncQuery(uid int32) (*ptttype.UserecRaw, error) {
	user, err := cmbbs.PasswdQuery(uid)
	if err != nil {
		return nil, err
	}

	user.Money = cache.MoneyOf(uid)

	return user, nil
}
