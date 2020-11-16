package cmbbs

import (
	"encoding/binary"
	"os"
	"reflect"

	"github.com/PichuChen/go-bbs/crypt"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
)

//CheckPasswd
//Params
//	expected: expected-passwd-hash
//	input: input-passwd
//
//Return
//	bool: true: good (password matched). false: bad (password not matched).
//	error: err
func CheckPasswd(expected []byte, input []byte) (bool, error) {
	pw, err := crypt.Fcrypt(input, expected)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(pw, expected), nil
}

func LogAttempt(userID *[ptttype.IDLEN + 1]byte, ip *[ptttype.IPV4LEN + 1]byte, isWithUserHome bool) {
}

//PasswdLoadUser
//Params
//	userID: user-id
//
//Return
//	int: user-num in passwd file.
//	*ptttype.UserecRaw: user.
//	error: err.
func PasswdLoadUser(userID *[ptttype.IDLEN + 1]byte) (int, *ptttype.UserecRaw, error) {
	if userID == nil || userID[0] == 0 {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	usernum, _, err := shm.SearchUserRaw(userID[:], false)
	if err != nil {
		return 0, nil, err
	}

	if usernum > ptttype.MAX_USERS {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	user, err := PasswdQuery(usernum)
	if err != nil {
		return 0, nil, err
	}

	return usernum, user, nil
}

//PasswdQuery
//Params
//	num: user-num in passwd file.
//
//Return
//	*ptttype.UserecRaw: user.
//	error: err.
func PasswdQuery(num int) (*ptttype.UserecRaw, error) {
	if num < 1 || num > ptttype.MAX_USERS {
		return nil, ptttype.ErrInvalidUserID
	}

	file, err := os.Open(ptttype.FN_PASSWD)
	if err != nil {
		return nil, err
	}

	user := &ptttype.UserecRaw{}
	offset := ptttype.USEREC_RAW_SZ * (int64(num) - 1)
	_, err = file.Seek(offset, 0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}