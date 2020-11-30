package ptttype

type RestrictReason uint8

const (
	_ RestrictReason = iota
	RESTRICT_REASON_FORBIDDEN
	RESTRICT_REASON_HIDDEN
)

type BoardSummary struct {
	Bid     int32
	Attr    BoardStatAttr
	Brdname [IDLEN + 1]byte
	Title   []byte
	BM      [IDLEN*3 + 3]byte
	Reason  RestrictReason
}
