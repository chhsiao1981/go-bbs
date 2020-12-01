package ptttype

import "github.com/PichuChen/go-bbs/types"

type RestrictReason uint8

const (
	_ RestrictReason = iota
	RESTRICT_REASON_FORBIDDEN
	RESTRICT_REASON_HIDDEN
)

type BoardSummary struct {
	Bid          int32
	BrdAttr      BrdAttr
	StatAttr     BoardStatAttr
	Brdname      BoardID_t
	RealTitle    []byte
	BM           []UserID_t
	Reason       RestrictReason
	LastPostTime types.Time4
	NUser        int32
	Total        int32
}
