package ptttype

type RestrictReason uint8

const (
	_ RestrictReason = iota
	RESTRICT_REASON_FORBIDDEN
	RESTRICT_REASON_HIDDEN
)

type BoardSummary struct {
	Bid       int32
	Attr      BrdAttr
	StatAttr  BoardStatAttr
	Brdname   BoardID_t
	RealTitle []byte
	BM        BM_t
	Reason    RestrictReason
}
