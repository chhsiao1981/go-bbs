package ptttype

import "unsafe"

var (
	EMPTY_USER_ID  = [IDLEN + 1]byte{}
	EMPTY_BOARD_ID = [IDLEN + 1]byte{}
)

const USER_ID_SZ = unsafe.Sizeof(EMPTY_USER_ID)
const BOARD_ID_SZ = unsafe.Sizeof(EMPTY_BOARD_ID)
