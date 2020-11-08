package bbs

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

type GetBoardInfo struct {
	brdname string // bid
	title   string
	BM      string // bms uid
	brdattr uint32
	level   uint32
	nuser   int32
}
