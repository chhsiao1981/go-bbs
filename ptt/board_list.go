package ptt

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//LoadGeneralBoards
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1142
func LoadGeneralBoards(user *ptttype.UserecRaw, uid int32, startBID int32, nBoards int32, keyword []byte) (summary []*ptttype.BoardSummary, nextBID int32, err error) {

	nBoardsInCache := cache.NumBoards()

	endBID := startBID + nBoards // endBID is exluded
	if endBID >= nBoardsInCache {
		endBID = nBoardsInCache
	}

	boardStats := make([]*ptttype.BoardStat, 0, endBID-startBID)
	for i := startBID; i < endBID; i++ {
		eachBoardStat, err := loadGeneralBoardStat(user, uid, i, keyword)
		if err != nil {
			continue
		}
		if eachBoardStat == nil {
			continue
		}

		boardStats = append(boardStats, eachBoardStat)
	}

	summary, err = showBoardList(user, uid, boardStats)
	if err != nil {
		return nil, -1, err
	}

	if endBID == nBoardsInCache {
		endBID = -1
	}

	return summary, endBID, nil
}

//loadGeneralBoardStat
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1147
func loadGeneralBoardStat(user *ptttype.UserecRaw, uid int32, idx int32, keyword []byte) (*ptttype.BoardStat, error) {
	var bidInCache int32

	const bsort0sz = unsafe.Sizeof(cache.Shm.Raw.BSorted[0])
	cache.Shm.ReadAt(
		unsafe.Offsetof(cache.Shm.Raw.BSorted)+bsort0sz*uintptr(ptttype.BSORT_BY_GENERAL)+uintptr(idx)*types.INT32_SZ,
		types.INT32_SZ,
		unsafe.Pointer(&bidInCache),
	)
	if bidInCache < 0 {
		return nil, nil
	}

	board := &ptttype.BoardHeaderRaw{}
	cache.Shm.ReadAt(
		unsafe.Offsetof(cache.Shm.Raw.BCache)+ptttype.BOARD_HEADER_RAW_SZ*uintptr(bidInCache),
		ptttype.BOARD_HEADER_RAW_SZ,
		unsafe.Pointer(board),
	)

	bid := bidInCache + 1
	isGroupOp := groupOp(user, board)
	state := boardPermStat(user, uid, board, bid)
	if (board.Brdname[0] == '\x00') ||
		(board.BrdAttr&(ptttype.BRD_GROUPBOARD|ptttype.BRD_SYMBOLIC) != 0) ||
		!((state != ptttype.NBRD_INVALID) || isGroupOp) ||
		keywordNotInTitle(&board.Title, keyword) {
		return nil, nil
	}

	boardStat := newBoardStat(bidInCache, state, board, isGroupOp)

	return boardStat, nil
}

//newBoardStat
func newBoardStat(bidInCache int32, state ptttype.BoardStatAttr, board *ptttype.BoardHeaderRaw, isGroupOp bool) (boardStat *ptttype.BoardStat) {
	boardStat = &ptttype.BoardStat{}

	boardStat.Bid = bidInCache + 1
	boardStat.Attr = state

	boardStat.Board = board
	boardStat.IsGroupOp = isGroupOp

	//XXX need to modify this by having state with NBRD_SET_POSTMASK
	//XXX this is a hack to ensure the brd-postmask
	var brd_postmask = ptttype.BRD_POSTMASK
	if (board.BrdAttr&ptttype.BRD_HIDE != 0) && (board.BrdAttr&ptttype.BRD_POSTMASK == 0) && state == ptttype.NBRD_BOARD {
		cache.Shm.SetOrUint32(
			unsafe.Offsetof(cache.Shm.Raw.BCache)+ptttype.BOARD_HEADER_RAW_SZ*uintptr(bidInCache)+ptttype.BOARD_HEADER_BRD_ATTR_OFFSET,
			unsafe.Pointer(&brd_postmask),
		)
		board.BrdAttr |= brd_postmask
	}

	return boardStat
}

//keywordNotInTitle
//
//TITLE_MATCH in board.c
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L14
func keywordNotInTitle(title *ptttype.BoardTitle_t, keyword []byte) bool {
	return (len(keyword) > 0) && (types.Cstrcasestr(title[:], keyword) < 0)
}

//showBoardList
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1409
func showBoardList(user *ptttype.UserecRaw, uid int32, boardStats []*ptttype.BoardStat) (summary []*ptttype.BoardSummary, err error) {
	summary = make([]*ptttype.BoardSummary, len(boardStats))
	for idx, eachStat := range boardStats {
		summary[idx] = parseBoardSummary(user, uid, eachStat)
	}

	return summary, nil
}

//parseBoardSummary
//
//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L1460
func parseBoardSummary(user *ptttype.UserecRaw, uid int32, boardStat *ptttype.BoardStat) (summary *ptttype.BoardSummary) {

	//XXX we do not deal with fav in go-bbs.
	if boardStat.Attr&ptttype.NBRD_LINE != 0 {
		return &ptttype.BoardSummary{Bid: boardStat.Bid, StatAttr: boardStat.Attr}
	}

	//XXX we do not deal with fav in go-bbs.
	if boardStat.Attr&ptttype.NBRD_FOLDER != 0 {
		return &ptttype.BoardSummary{Bid: boardStat.Bid, StatAttr: boardStat.Attr}
	}

	//hidden board
	board := boardStat.Board
	if !boardStat.IsGroupOp && boardStat.Attr == ptttype.NBRD_INVALID {
		reason := ptttype.RESTRICT_REASON_FORBIDDEN
		if board.BrdAttr&ptttype.BRD_HIDE != 0 {
			reason = ptttype.RESTRICT_REASON_HIDDEN
		}
		summary = &ptttype.BoardSummary{
			Bid:      boardStat.Bid,
			BrdAttr:  board.BrdAttr,
			StatAttr: boardStat.Attr,
			Brdname:  board.Brdname,
			Reason:   reason,
		}
		if ptttype.USE_REAL_DESC_FOR_HIDDEN_BOARD_IN_MYFAV {
			summary.RealTitle = board.Title.RealTitle()
		}

		return summary
	}

	bidInCache := boardStat.Bid - 1
	var lastPostTime types.Time4
	cache.Shm.ReadAt(
		unsafe.Offsetof(cache.Shm.Raw.LastPostTime)+types.TIME4_SZ*uintptr(bidInCache),
		types.TIME4_SZ,
		unsafe.Pointer(&lastPostTime),
	)

	var total int32
	cache.Shm.ReadAt(
		unsafe.Offsetof(cache.Shm.Raw.Total)+types.INT32_SZ*uintptr(bidInCache),
		types.INT32_SZ,
		unsafe.Pointer(&total),
	)

	summary = &ptttype.BoardSummary{
		Bid:          boardStat.Bid,
		BrdAttr:      board.BrdAttr,
		StatAttr:     boardStat.Attr,
		Brdname:      board.Brdname,
		BM:           board.BM.ToBMs(),
		LastPostTime: lastPostTime,
		NUser:        board.NUser,
		Total:        total,
	}

	return summary
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
