package names

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//IsValidUserID
//
//Params
//	userID: user-id
//
//Return
//	bool: is valid user-id
func IsValidUserID(userID *[ptttype.IDLEN + 1]byte) bool {
	if userID == nil {
		return false
	}

	theLen := types.Cstrlen(userID[:])
	if theLen < 2 || theLen > ptttype.IDLEN {
		return false
	}

	if !isalpha(userID[0]) {
		return false
	}

	for idx, c := range userID {
		if idx == theLen {
			break
		}

		if !isalnum(c) {
			return false
		}
	}

	return true
}

func isalpha(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}

func isnumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isalnum(c byte) bool {
	return isalpha(c) || isnumber(c)
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}
