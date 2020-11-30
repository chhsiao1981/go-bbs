package ptttype

type BoardStatAttr uint8

const (
	NBRD_FAV      BoardStatAttr = 1
	NBRD_BOARD    BoardStatAttr = 2
	NBRD_LINE     BoardStatAttr = 4
	NBRD_FOLDER   BoardStatAttr = 8
	NBRD_TAG      BoardStatAttr = 16
	NBRD_UNREAD   BoardStatAttr = 32
	NBRD_SYMBOLIC BoardStatAttr = 64
)

type BoardStat struct {
	Bid  int32
	Attr BoardStatAttr
}
