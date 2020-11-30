package path

import (
	"os"
	"strings"

	"github.com/PichuChen/go-bbs/names"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func SetHomePath(userID *[ptttype.IDLEN + 1]byte) string {
	return strings.Join([]string{
		ptttype.BBSHOME,
		ptttype.DIR_HOME,
		string(userID[0]),
		types.CstrToString(userID[:]),
	},
		string(os.PathSeparator),
	)
}

func SetHomeFile(userID *[ptttype.IDLEN + 1]byte, filename string) (string, error) {
	if !names.IsValidUserID(userID) {
		return "", ptttype.ErrInvalidUserID
	}
	if filename[0] == '\x00' || !IsValidFilename(filename) {
		return "", ptttype.ErrInvalidFilename
	}
	return strings.Join([]string{
		ptttype.BBSHOME,
		ptttype.DIR_HOME,
		string(userID[0]),
		types.CstrToString(userID[:]),
		filename,
	},
		string(os.PathSeparator),
	), nil
}

func IsValidFilename(filename string) bool {
	return !strings.Contains(filename, "..")
}

func SetBFile(boardID *[ptttype.IDLEN + 1]byte, filename string) (string, error) {
	if filename[0] == '\x00' || !IsValidFilename(filename) {
		return "", ptttype.ErrInvalidFilename
	}

	return strings.Join([]string{
		ptttype.BBSHOME,
		ptttype.DIR_BOARD,
		string(boardID[0]),
		types.CstrToString(boardID[:]),
		filename,
	},
		string(os.PathSeparator),
	), nil
}
