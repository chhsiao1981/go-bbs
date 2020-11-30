package ptttype

import (
	"encoding/binary"
	"io"
	"os"
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
)

type BoardHeaderRaw struct {
	//Require updating SHM_VERSION if BOARD_HEADER_RAW_SZ is changed.
	Brdname            BoardID_t /* bid */
	Title              BoardTitle_t
	BM                 BM_t /* BMs' userid, token '/' */
	Pad1               [3]byte
	BrdAttr            BrdAttr     /* board的屬性 */
	ChessCountry       int8        /* 棋國 */
	VoteLimitPosts_    uint8       /* (已停用) 連署 : 文章篇數下限 */
	VoteLimitLogins    uint8       /* 連署 : 登入次數下限 */
	Pad2_1             [1]uint8    /* (已停用) 連署 : 註冊時間限制 */
	BUpdate            types.Time4 /* note update time */
	PostLimitPosts_    uint8       /* (已停用) 發表文章 : 文章篇數下限 */
	PostLimitLogins    uint8       /* 發表文章 : 登入次數下限 */
	Pad2_2             [1]uint8    /* (已停用) 發表文章 : 註冊時間限制 */
	BVote              uint8       /* 正舉辦 Vote 數 */
	VTime              types.Time4 /* Vote close time */
	Level              PERM        /* 可以看此板的權限 */
	PermReload         types.Time4 /* 最後設定看板的時間 */
	Gid                int32       /* 看板所屬的類別 ID */
	Next               [2]int32    /* 在同一個gid下一個看板 動態產生*/
	FirstChild         [2]int32    /* 屬於這個看板的第一個子看板 */
	Parent             int32       /* 這個看板的 parent 看板 bid */
	ChildCount         int32       /* 有多少個child */
	NUser              int32       /* 多少人在這板 */
	PostExpire         int32       /* postexpire */
	EndGamble          types.Time4
	PostType           [33]byte
	PostTypeF          byte
	FastRecommendPause uint8 /* 快速連推間隔 */
	VoteLimitBadpost   uint8 /* 連署 : 劣文上限 */
	PostLimitBadpost   uint8 /* 發表文章 : 劣文上限 */
	Pad3               [3]byte
	SRexpire           types.Time4 /* SR Records expire time */
	Pad4               [40]byte
}

//!!!Require updating SHM_VERSION if BOARD_HEADER_RAW_SZ is changed.
var emptyBoardHeaderRaw = BoardHeaderRaw{}

const BOARD_HEADER_RAW_SZ = unsafe.Sizeof(emptyBoardHeaderRaw)

const BOARD_HEADER_BOARD_NAME_OFFSET = unsafe.Offsetof(emptyBoardHeaderRaw.Brdname)

func LoadBoardHeaderRawsFromFile(filename string) ([]*BoardHeaderRaw, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var boards []*BoardHeaderRaw
	for {
		eachBoard, err := NewBoardHeaderRawWithFile(file)
		if err != nil {
			if err == io.EOF {
				return boards, nil
			} else {
				return nil, err
			}
		}
		boards = append(boards, eachBoard)
	}
}

func NewBoardHeaderRawWithFile(file *os.File) (*BoardHeaderRaw, error) {
	boardRaw := &BoardHeaderRaw{}

	err := binary.Read(file, binary.LittleEndian, boardRaw)
	if err != nil {
		return nil, err
	}

	return boardRaw, nil
}
