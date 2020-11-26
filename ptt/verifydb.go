package ptt

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//verifyDBSet
//
//XXX to implement
func verifyDBSet(userID *[ptttype.IDLEN + 1]byte, generation int64, vmethod ptttype.VerifyDBVMethod, vkey []byte, timestamp int64) error {
	return types.ErrNotImplemented
}

//verifyDBCountByVerify
//
//XXX to implement
func verifyDBCountByVerify(vmethod ptttype.VerifyDBVMethod, vkey []byte) (countSelf int, countOther int, err error) {
	return 0, 0, types.ErrNotImplemented
}
