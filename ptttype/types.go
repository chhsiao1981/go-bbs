package ptttype

import "unsafe"

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
