package ptttype

import (
	"bytes"
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
)

//A good thing about specifying the type,
//is that we don't need to worry whether we need to + 1 or not.
//We don't mix-up with user-id and board-id either.

type UserID_t [IDLEN + 1]byte
type RealName_t [REALNAMESZ]byte
type Nickname_t [NICKNAMESZ]byte
type Passwd_t [PASSLEN]byte
type IPv4_t [IPV4LEN + 1]byte
type Email_t [EMAILSZ]byte
type Address_t [ADDRESSSZ]byte
type Reg_t [REGLEN + 1]byte
type Career_t [CAREERSZ]byte
type Phone_t [PHONESZ]byte

type BoardID_t [IDLEN + 1]byte
type BM_t [IDLEN*3 + 3]byte

type Filename_t [FNLEN]byte
type Subject_t [STRLEN]byte
type RCPT_t [RCPTSZ]byte

type Owner_t [IDLEN + 2]byte //uid[.]

type Title_t [TTLEN + 1]byte

type Date_t [6]byte

var (
	EMPTY_USER_ID  = UserID_t{}
	EMPTY_BOARD_ID = BoardID_t{}
	EMPTY_EMAIL    = Email_t{}
)

const USER_ID_SZ = unsafe.Sizeof(EMPTY_USER_ID)
const BOARD_ID_SZ = unsafe.Sizeof(EMPTY_BOARD_ID)
const EMAIL_SZ = unsafe.Sizeof(EMPTY_EMAIL)

type BoardTitle_t [BTLEN + 1]byte

func (bt *BoardTitle_t) RealTitle() []byte {
	return bt[7:]
}

//ToBMs
//
//We would like to have a better method
//(We don't need to worry about this once we move everything to the db.)
func (bm *BM_t) ToBMs() []UserID_t {
	bmBytes := types.CstrToBytes(bm[:])
	theList := bytes.Split(bmBytes, []byte{'/'})
	bms := make([]UserID_t, len(theList))
	for idx, each := range theList {
		copy(bms[idx][:], each[:])
	}

	return bms
}
