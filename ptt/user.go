package ptt

import (
	"os"
	"strings"

	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

//killUser
//
//Assume correct uid / userID correspondance.
func killUser(uid int32, userID *[ptttype.IDLEN + 1]byte) error {
	if uid <= 0 || userID == nil {
		return ptttype.ErrInvalidUserID
	}

	err := friendDeleteAll(userID, ptttype.FRIEND_ALOHA)
	if err != nil {
		log.Errorf("killUser: unable to friend-delete-all: uid: %v e: %v", uid, err)
	}

	err = tryDeleteHomePath(userID)
	if err != nil {
		log.Errorf("killUser: unable to delete home-path: userID: %v e: %v", userID, err)
	}

	emptyUser := &ptttype.UserecRaw{}
	err = passwdSyncUpdate(uid, emptyUser)
	if err != nil {
		log.Errorf("killUser: unable to passwd-sync-update emptyUser: uid: %v e: %v", uid, err)
	}

	return nil
}

func tryDeleteHomePath(userID *[ptttype.IDLEN + 1]byte) error {
	homePath := cmbbs.SetHomePath(userID)
	dstPath := strings.Join([]string{ptttype.BBSHOME, ptttype.DIR_TMP, types.CstrToString(userID[:])}, string(os.PathSeparator))

	if !types.IsDir(homePath) {
		return nil
	}

	if err := types.Rename(homePath, dstPath); err != nil {
		return err
	}

	if err := os.RemoveAll(homePath); err != nil {
		return err
	}

	return nil
}
