package ptt

import (
	"bytes"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func is_uBM(userID *ptttype.UserID_t, bm *ptttype.BM_t) bool {
	userIDBytes := types.CstrToBytes(userID[:])
	bmBytes := types.CstrToBytes(bm[:])
	theIdx := bytes.Index(bmBytes, userIDBytes)
	if theIdx < 0 {
		return false
	}

	isValidHead := true
	if theIdx > 0 {
		isValidHead = !types.Isalnum(bmBytes[theIdx-1])
	}

	isValidTail := true
	if theIdx+len(userIDBytes) < len(bmBytes) {
		isValidTail = !types.Isalnum(bmBytes[theIdx+len(userIDBytes)])
	}

	return isValidHead && isValidTail
}
