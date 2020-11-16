package cache

import (
	"encoding/binary"
	"os"
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
)

//passwdUpdateMoney
//
//XXX should not call this directly.
//    call this from DeUMoney (SetUMoney).
func passwdUpdateMoney(uid int32, money int32) error {
	if uid < 1 || uid >= ptttype.MAX_USERS {
		return ErrInvalidUID
	}

	file, err := os.OpenFile(ptttype.FN_PASSWD, os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	const offsetMoney = unsafe.Offsetof(ptttype.USEREC_RAW.Money)
	file.Seek(int64(ptttype.USEREC_RAW_SZ*uintptr(uid-1)+offsetMoney), 0)
	binary.Write(file, binary.LittleEndian, &money)

	return nil
}
