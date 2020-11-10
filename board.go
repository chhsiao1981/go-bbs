package bbs

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

type BoardHeader struct {
	BrdName            string
	Title              string
	BM                 []string
	Brdattr            ptttype.BrdAttr // uid[.]
	ChessCountry       int8
	VoteLimitPosts     uint8
	VoteLimitLogins    uint8
	BUpdate            time.Time
	PostLimitPosts     uint8
	PostLimitLogins    uint8
	BVote              uint8
	VTime              time.Time
	Level              ptttype.PERM
	PermReload         time.Time
	Gid                int32
	Next               []int32
	FirstChild         []int32
	Parent             int32
	ChildCount         int32
	Nuser              int32
	PostExpire         int32
	EndGamble          time.Time
	PostType           []byte
	PostTypeF          byte
	FastRecommendPause uint8
	VoteLimitBadPost   uint8
	PostLimitBadPost   uint8
	SRexpire           time.Time
}

const (
	PTT_BRD_POSTMASK   = 0x00000020
	PTT_BRD_GROUPBOARD = 0x00000008
	PTT_PERM_SYSOP     = 000000040000
	PTT_PERM_BM        = 000000002000
	PTT_BRD_HIDE       = 0x00000010
)

func NewBoardHeaderFromRaw(boardRaw *ptttype.BoardHeaderRaw) *BoardHeader {
	board := &BoardHeader{}
	board.BrdName = Big5ToUtf8(types.CstrToBytes(boardRaw.Brdname[:]))
	board.Title = Big5ToUtf8(types.CstrToBytes(boardRaw.Title[:]))

	utf8Str := Big5ToUtf8(types.CstrToBytes(boardRaw.BM[:]))
	if len(utf8Str) == 0 {
		board.BM = make([]string, 0)
	} else {
		board.BM = strings.Split(utf8Str, string("/"))
	}
	board.Brdattr = boardRaw.BrdAttr
	board.ChessCountry = boardRaw.ChessCountry
	board.VoteLimitPosts = boardRaw.VoteLimitPosts_
	board.VoteLimitLogins = boardRaw.VoteLimitLogins
	board.BUpdate = boardRaw.BUpdate.ToLocal()
	board.PostLimitPosts = boardRaw.PostLimitPosts_
	board.PostLimitLogins = boardRaw.PostLimitLogins
	board.BVote = boardRaw.BVote
	board.VTime = boardRaw.VTime.ToLocal()
	board.Level = boardRaw.Level
	board.PermReload = boardRaw.PermReload.ToLocal()
	board.Gid = boardRaw.Gid
	board.Next = boardRaw.Next[:]
	board.FirstChild = boardRaw.FirstChild[:]
	board.Parent = boardRaw.Parent
	board.ChildCount = boardRaw.ChildCount
	board.Nuser = boardRaw.NUser
	board.PostExpire = boardRaw.PostExpire
	board.EndGamble = boardRaw.EndGamble.ToLocal()
	board.PostType = boardRaw.PostType[:]
	board.PostTypeF = boardRaw.PostTypeF
	board.FastRecommendPause = boardRaw.FastRecommendPause
	board.VoteLimitBadPost = boardRaw.VoteLimitBadpost
	board.PostLimitBadPost = boardRaw.PostLimitBadpost
	board.SRexpire = boardRaw.SRexpire.ToLocal()

	return board
}

func OpenBoardHeaderFile(filename string) ([]*BoardHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	ret := []*BoardHeader{}

	for {
		board, eachErr := NewBoardHeaderWithFile(file)
		if eachErr != nil {
			// io.EOF is reading correctly to the end the file.
			if eachErr == io.EOF {
				break
			}

			err = eachErr
			break
		}
		ret = append(ret, board)
	}

	return ret, err

}

func NewBoardHeaderWithFile(file *os.File) (*BoardHeader, error) {
	boardRaw, err := ptttype.NewBoardHeaderRawWithFile(file)
	if err != nil {
		return nil, err
	}

	user := NewBoardHeaderFromRaw(boardRaw)

	return user, nil
}
