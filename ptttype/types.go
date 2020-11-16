package ptttype

import "unsafe"

const USER_ID_SZ = unsafe.Sizeof([IDLEN + 1]byte{})
const BOARD_ID_SZ = unsafe.Sizeof([IDLEN + 1]byte{})
