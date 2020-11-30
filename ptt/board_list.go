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
		return &ptttype.BoardSummary{Bid: boardStat.Bid, Attr: boardStat.Attr}, nil
	}

	//XXX we do not deal with fav in go-bbs.
	if boardStat.Attr&ptttype.NBRD_FOLDER != 0 {
		return &ptttype.BoardSummary{Bid: boardStat.Bid, Attr: boardStat.Attr}, nil
	}

	//hidden board
	board, err := cache.GetBCache(boardStat.Bid - 1)
	if err != nil {
		return nil, err
	}
	if !groupOp() && !hasBoardPerm(user, uid, board, boardStat.Bid) {
		reason := ptttype.RESTRICT_REASON_FORBIDDEN
		if board.BrdAttr&ptttype.BRD_HIDE != 0 {
			reason = ptttype.RESTRICT_REASON_HIDDEN
		}
		summary = &ptttype.BoardSummary{
			Bid:     boardStat.Bid,
			Attr:    boardStat.Attr,
			Brdname: board.Brdname,
			Reason:  reason,
		}
		if ptttype.USE_REAL_DESC_FOR_HIDDEN_BOARD_IN_MYFAV {
			summary.Title = board.Title.RealTitle()
		}

		return summary, nil
	}
	return nil, nil
}

func groupOp() bool {
	//XXX TODO: implement groupOp
	return true
}
