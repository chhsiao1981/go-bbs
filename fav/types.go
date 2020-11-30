package fav

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
)

type FAVH int8

const (
	FAVH_FAV     FAVH = 1
	FAVH_TAG     FAVH = 2
	FAVH_UNREAD  FAVH = 4
	FAVH_ADM_TAG FAVH = 8
)

type FavLine struct {
	Lid int8
}

const SIZE_OF_FAV_LINE = unsafe.Sizeof(FavLine{})

type FavFolder struct {
	Fid        int8
	Title      [ptttype.BTLEN + 1]byte
	ThisFolder *FavRaw
}

const SIZE_OF_FAV_FOLDER = unsafe.Sizeof(FavFolder{})

type FavBoard struct {
	Bid       int32
	LastVisit int32 /* UNUSED */
	Attr      int8
}

const SIZE_OF_FAV_BOARD = unsafe.Sizeof(FavBoard{})

type Fav4Folder struct {
	Fid        int8
	Title      [ptttype.BTLEN + 1]byte
	ThisFolder int32
}

const SIZE_OF_FAV4_FOLDER = unsafe.Sizeof(Fav4Folder{})

type Fav4Board struct {
	Bid       int32
	LastVisit int32
	Attr      int8
}

const SIZE_OF_FAV4_BOARD = unsafe.Sizeof(Fav4Board{})
