package ptt

import (
	"bytes"

	"github.com/PichuChen/go-bbs/ptttype"
)

func pwcuEnableBit(perm ptttype.PERM, mask ptttype.PERM) ptttype.PERM {
	return perm | mask
}

func pwcuDisableBit(perm ptttype.PERM, mask ptttype.PERM) ptttype.PERM {
	return perm & ^mask
}

func pwcuSetByBit(perm ptttype.PERM, mask ptttype.PERM, isSet bool) ptttype.PERM {
	if isSet {
		return pwcuEnableBit(perm, mask)
	} else {
		return pwcuDisableBit(perm, mask)
	}
}

func pwcuStart(uid int32, userID *[ptttype.IDLEN + 1]byte) (user *ptttype.UserecRaw, err error) {
	user, err = passwdSyncQuery(uid)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(userID[:], user.UserID[:]) {
		return nil, ptttype.ErrInvalidUserID
	}

	return user, nil
}

func pwcuEnd(uid int32, user *ptttype.UserecRaw) (err error) {
	return passwdSyncUpdate(uid, user)
}

func pwcuRegCompleteJustify(uid int32, userID *[ptttype.IDLEN + 1]byte, justify *[ptttype.REGLEN + 1]byte) (err error) {

	user, err := pwcuStart(uid, userID)
	if err != nil {
		return err
	}

	copy(user.Justify[:], justify[:])
	user.UserLevel = pwcuEnableBit(user.UserLevel, ptttype.PERM_POST|ptttype.PERM_LOGINOK)

	err = pwcuEnd(uid, user)
	if err != nil {
		return err
	}

	return nil
}

func pwcuBitDisableLevel(uid int32, userID *[ptttype.IDLEN + 1]byte, perm ptttype.PERM) (err error) {
	user, err := pwcuStart(uid, userID)
	if err != nil {
		return err
	}

	pwcuDisableBit(user.UserLevel, perm)

	return pwcuEnd(uid, user)
}
