package cmbbs

import "C"
import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/rand"
	"os"

	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/crypt"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/sem"
)

//GenPasswd
//
//If passwd as empty: return empty passwd (unable to login)
func GenPasswd(passwd []byte) (passwdHash *ptttype.Passwd_t, err error) {
	if passwd[0] == 0 {
		return &ptttype.Passwd_t{}, nil
	}

	num := rand.Intn(65536)
	saltc := [2]byte{
		byte(num & 0x7f),
		byte((num >> 8) & 0x7f),
	}

	result, err := crypt.Fcrypt(passwd, saltc[:])
	passwdHash = &ptttype.Passwd_t{}
	copy(passwdHash[:], result[:])
	return passwdHash, err
}

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
	return bytes.Equal(pw, expected), nil //requires the passwd-hash be exact match.
}

func LogAttempt(userID *ptttype.UserID_t, ip *ptttype.IPv4_t, isWithUserHome bool) {
}

//PasswdLoadUser
//Params
//	userID: user-id
//
//Return
//	int32: uid
//	*ptttype.UserecRaw: user.
//	error: err.
func PasswdLoadUser(userID *ptttype.UserID_t) (int32, *ptttype.UserecRaw, error) {
	if userID == nil || userID[0] == 0 {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	uid, err := cache.SearchUserRaw(userID, nil)
	if err != nil {
		return 0, nil, err
	}

	if uid > ptttype.MAX_USERS {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	user, err := PasswdQuery(uid)
	if err != nil {
		return 0, nil, err
	}

	return uid, user, nil
}

//PasswdQuery
//Params
//	uid: uid
//
//Return
//	*ptttype.UserecRaw: user.
//	error: err.
func PasswdQuery(uid int32) (*ptttype.UserecRaw, error) {
	if uid < 1 || uid > ptttype.MAX_USERS {
		return nil, ptttype.ErrInvalidUserID
	}

	file, err := os.Open(ptttype.FN_PASSWD)
	if err != nil {
		return nil, err
	}

	user := &ptttype.UserecRaw{}
	offset := int64(ptttype.USEREC_RAW_SZ) * (int64(uid) - 1)
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

func PasswdUpdate(uid int32, user *ptttype.UserecRaw) error {
	if uid < 1 || uid > ptttype.MAX_USERS {
		return cache.ErrInvalidUID
	}

	file, err := os.OpenFile(ptttype.FN_PASSWD, os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(int64(ptttype.USEREC_RAW_SZ)*int64(uid-1), 0)
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.LittleEndian, user)
	if err != nil {
		return err
	}

	return nil
}

func PasswdInit() error {
	if Sem != nil {
		return nil
	}

	var err error
	Sem, err = sem.SemGet(ptttype.PASSWDSEM_KEY, 1, sem.SEM_R|sem.SEM_A|sem.IPC_CREAT|sem.IPC_EXCL)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			Sem, err = sem.SemGet(ptttype.PASSWDSEM_KEY, 1, sem.SEM_R|sem.SEM_A)
			if err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	err = Sem.SetVal(0, 1)
	if err != nil {
		return err
	}
	return nil
}

func PasswdLock() error {
	return Sem.Wait(0)
}

func PasswdUnlock() error {
	return Sem.Post(0)
}
