package ptt

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
)

//showBoardList
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1409
func showBoardList(user *ptttype.UserecRaw, uid int32, boardStats []*ptttype.BoardStat) (summary []*ptttype.BoardSummary, err error) {
	summary = make([]*ptttype.BoardSummary, len(boardStats))
	for idx, eachStat := range boardStats {
		eachSummary, err := parseBoardSummary(user, uid, eachStat)
		if err != nil {
			return nil, err
		}
		summary[idx] = eachSummary
	}

	return summary, nil
}

func parseBoardSummary(user *ptttype.UserecRaw, uid int32, boardStat *ptttype.BoardStat) (summary *ptttype.BoardSummary, err error) {

	//XXX we do not deal with fav in go-bbs.
	if boardStat.Attr&ptttype.NBRD_LINE != 0 {
		return &ptttype.BoardSummary{Bid: boardStat.Bid, StatAttr: boardStat.Attr}, nil
	}

	//XXX we do not deal with fav in go-bbs.
	if boardStat.Attr&ptttype.NBRD_FOLDER != 0 {
		return &ptttype.BoardSummary{Bid: boardStat.Bid, StatAttr: boardStat.Attr}, nil
	}

	//hidden board
	board, err := cache.GetBCache(boardStat.Bid - 1)
	if err != nil {
		return nil, err
	}
	if !groupOp(user, board) && !hasBoardPerm(user, uid, board, boardStat.Bid) {
		reason := ptttype.RESTRICT_REASON_FORBIDDEN
		if board.BrdAttr&ptttype.BRD_HIDE != 0 {
			reason = ptttype.RESTRICT_REASON_HIDDEN
		}
		summary = &ptttype.BoardSummary{
			Bid:      boardStat.Bid,
			Attr:     board.BrdAttr,
			StatAttr: boardStat.Attr,
			Brdname:  board.Brdname,
			Reason:   reason,
		}
		if ptttype.USE_REAL_DESC_FOR_HIDDEN_BOARD_IN_MYFAV {
			summary.RealTitle = board.Title.RealTitle()
		}

		return summary, nil
	}
	return nil, nil
}

//groupOp
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1579
func groupOp(user *ptttype.UserecRaw, board *ptttype.BoardHeaderRaw) bool {
	if hasUserPerm(user, ptttype.PERM_NOCITIZEN) {
		return false
	}

	if hasUserPerm(user, ptttype.PERM_BOARD) {
		return true
	}

	if is_uBM(&user.UserID, &board.BM) {
		return true
	}

	return false
}
