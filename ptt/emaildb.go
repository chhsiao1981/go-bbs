package ptt

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//emailDBCheckEmail
//
//XXX to implement
func emailDBCheckEmail(userID *ptttype.UserID_t, email *[ptttype.EMAILSZ]byte) (count int, err error) {

	return -1, types.ErrNotImplemented
}

//emailDBUpdateEmail
//
//XXX to implemenet
func emailDBUpdateEmail(userID *[ptttype.IDLEN + 1]byte, email *[ptttype.EMAILSZ]byte) (err error) {
	return types.ErrNotImplemented
}
