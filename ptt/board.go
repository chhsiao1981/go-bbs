package ptt

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
)

func hasBoardPerm(user *ptttype.UserecRaw, uid int32, board *ptttype.BoardHeaderRaw, bid int32) bool {
	if hasUserPerm(user, ptttype.PERM_SYSOP) {
		return true
	}

	return hasBoardPermNormally(user, uid, board, bid)
}

func hasBoardPermNormally(user *ptttype.UserecRaw, uid int32, board *ptttype.BoardHeaderRaw, bid int32) bool {
	level := board.Level
	brdAttr := board.BrdAttr

	// allow POLICE to enter BM boards
	if (level&ptttype.PERM_BM != 0) && (hasUserPerm(user, ptttype.PERM_POLICE) || hasUserPerm(user, ptttype.PERM_POLICE_MAN)) {
		return true
	}

	/* 板主 */
	if cache.IsBMCache(user, bid) {
		return true
	}

	/* 祕密看板：核對首席板主的好友名單 */
	if brdAttr&ptttype.BRD_HIDE != 0 {
		if !cache.IsHiddenBoardFriend(bid-1, uid-1) {
			if brdAttr&ptttype.BRD_POSTMASK != 0 {
				return false
			} else {
				return true //XXX return 2; //what's this?
			}
		} else {
			return true
		}
	}

	// TODO Change this to a query on demand.
	/* 十八禁看板 */
	if brdAttr&ptttype.BRD_OVER18 != 0 && !user.Over18 {
		return false
	}

	if level != 0 && (brdAttr&ptttype.BRD_POSTMASK) == 0 && !hasUserPerm(user, level) {
		return false
	}

	return true
}
