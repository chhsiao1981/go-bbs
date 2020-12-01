package ptt

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
)

//boardPermStat
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L197
//
//The original hashBoardPerm
func boardPermStat(user *ptttype.UserecRaw, uid int32, board *ptttype.BoardHeaderRaw, bid int32) ptttype.BoardStatAttr {

	//SYSOP
	if hasUserPerm(user, ptttype.PERM_SYSOP) {
		return ptttype.NBRD_FAV
	}

	return boardPermStatNormally(user, uid, board, bid)
}

//boardPermStatNormally
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L157
//
//The original hasBoardPermNormally
//
//Original code was mixing-up BoardStat with this function, making the code not easy to comprehend.
//BM / Police / SYSOP treat the board as NBRD_FAV, while others treat the board as NBRD_BOARD.
//It's because that newBoardStat is hacked to forcely add BRD_POSTMASK if not set properly and the type is NBRD_BOARD.
//Need to figure out a better method to solve this issue.
func boardPermStatNormally(user *ptttype.UserecRaw, uid int32, board *ptttype.BoardHeaderRaw, bid int32) ptttype.BoardStatAttr {
	level := board.Level
	brdAttr := board.BrdAttr

	// allow POLICE to enter BM boards
	if (level&ptttype.PERM_BM != 0) && (hasUserPerm(user, ptttype.PERM_POLICE) || hasUserPerm(user, ptttype.PERM_POLICE_MAN)) {
		return ptttype.NBRD_FAV
	}

	/* 板主 */
	if cache.IsBMCache(user, bid) {
		return ptttype.NBRD_FAV
	}

	/* 祕密看板：核對首席板主的好友名單 */
	if brdAttr&ptttype.BRD_HIDE != 0 {
		if !cache.IsHiddenBoardFriend(bid-1, uid-1) {
			if brdAttr&ptttype.BRD_POSTMASK != 0 {
				return ptttype.NBRD_INVALID
			} else {
				return ptttype.NBRD_BOARD //XXX return 2; //what's this? (in addnewbrdstat, to set brd_postmask)
			}
		} else {
			return ptttype.NBRD_FAV
		}
	}

	// TODO Change this to a query on demand.
	/* 十八禁看板 */
	if brdAttr&ptttype.BRD_OVER18 != 0 && !user.Over18 {
		return ptttype.NBRD_INVALID
	}

	if level != 0 && (brdAttr&ptttype.BRD_POSTMASK) == 0 && !hasUserPerm(user, level) {
		return ptttype.NBRD_INVALID
	}

	return ptttype.NBRD_FAV
}
