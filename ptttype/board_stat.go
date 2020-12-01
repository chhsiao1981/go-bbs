package ptttype

type BoardStatAttr uint8

const (
	NBRD_INVALID  BoardStatAttr = 0
	NBRD_FAV      BoardStatAttr = 1
	NBRD_BOARD    BoardStatAttr = 2
	NBRD_LINE     BoardStatAttr = 4
	NBRD_FOLDER   BoardStatAttr = 8
	NBRD_TAG      BoardStatAttr = 16
	NBRD_UNREAD   BoardStatAttr = 32
	NBRD_SYMBOLIC BoardStatAttr = 64
)

type BoardStat struct {
	//BoardStat
	//
	//    BoardStat should be used as read-only process,
	//    Or we need to be very careful that the referenced
	//    board does not write back to Shm or file.
	Bid  int32
	Attr BoardStatAttr

	Board *BoardHeaderRaw

	//obtain first in load-board-stat to be processed for showBrdList
	IsGroupOp bool
}
