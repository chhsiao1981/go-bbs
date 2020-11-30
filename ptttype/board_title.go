package ptttype

type BoardTitle [BTLEN + 1]byte

func (bt *BoardTitle) RealTitle() []byte {
	return bt[7:]
}
