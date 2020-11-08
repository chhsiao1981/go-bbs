package bbs

import (
	"encoding/binary"
	"log"
	"os"
	"unsafe"
)

const (
	FAV_VERSION = 3363
)

const (
	FAVT_BOARD  = 1
	FAVT_FOLDER = 2
	FAVT_LINE   = 3
)

const (
	FAVH_FAV     = 1
	FAVH_TAG     = 2
	FAVH_UNREAD  = 4
	FAVH_ADM_TAG = 8
)

const (
	FAV_PRE_ALLOC     = 8
	FAV_MAXDEPTH      = 5
	MAX_FAV           = 1024
	MAX_LINE          = 64
	MAX_FOLDER        = 64
	NEW_FAV_THRESHOLD = 12
)

const (
	FAV   = ".fav"
	FAV4  = ".fav4"
	FAVNB = ".favnb"
)

type FavType struct {
	TheType int8
	Attr    int8
	Fp      interface{}
}

type Fav struct {
	NAllocs  int16
	DataTail int16 /* the tail of item list that user
	   have ever used */
	NBoards  int16 /* number of the boards */
	NLines   int8  /* number of the lines */
	NFolders int8  /* number of the folders */
	LineID   int8  /* current max line id */
	FolderID int8  /* current max folder id */

	Favh []*FavType
}

const SIZE_OF_FAV = unsafe.Sizeof(Fav{})

type FavBoard struct {
	Bid       int32
	LastVisit int32 /* UNUSED */
	Attr      byte
}

const SIZE_OF_FAV_BOARD = unsafe.Sizeof(FavBoard{})

type FavFolder struct {
	Fid        int8
	TitleBig5  [BTLEN + 1]byte
	ThisFolder *Fav
}

const SIZE_OF_FAV_FOLDER = unsafe.Sizeof(FavFolder{})

type FavLine struct {
	Lid int8
}

const SIZE_OF_FAV_LINE = unsafe.Sizeof(FavLine{})

func FavLoad(userID string) (*Fav, error) {
	filename, err := readFavFilename(userID)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// version
	version := int16(0)
	binary.Read(file, binary.LittleEndian, &version)
	if version != FAV_VERSION {
		log.Println("version is not expected version")
	}

	// favrec
	favrec, err := readFavrec(file)
	if err != nil {
		return nil, err
	}

	return favrec, nil
}

func readFavFilename(userID string) (string, error) {
	favBoardFilename := setuserfile(userID, FAV)

	if !dashf(favBoardFilename) {

	}

	return favBoardFilename, nil
}

func readFavrec(file *os.File) (*Fav, error) {
	// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c
	fav := &Fav{}
	err := binary.Read(file, binary.LittleEndian, &fav.NBoards)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &fav.NLines)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &fav.NFolders)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	fav.DataTail = getDataNumber(fav)
	fav.NAllocs = fav.DataTail + FAV_PRE_ALLOC
	fav.LineID = 0
	fav.FolderID = 0
	fav.Favh = make([]*FavType, fav.NAllocs)

	for i := int16(0); i < fav.DataTail; i++ {
		ft := &FavType{}
		fav.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return nil, ErrInvalidFavType
		}
		if !isValidFavType(ft.TheType) {
			return nil, ErrInvalidFavType
		}

		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return nil, ErrInvalidFavType
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := castFolder(ft)
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, favFolder.TitleBig5)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
		case FAVT_BOARD:
			favBoard := castBoard(ft)
			if favBoard == nil {
				return nil, ErrInvalidFavBoard
			}
			err = binary.Read(file, binary.LittleEndian, favBoard)
			if err != nil {
				return nil, ErrInvalidFavBoard
			}
		case FAVT_LINE:
			favLine := castLine(ft)
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			err = binary.Read(file, binary.LittleEndian, favLine)
			if err != nil {
				return nil, ErrInvalidFavLine
			}
		}
	}

	for i := int16(0); i < fav.DataTail; i++ {
		ft := fav.Favh[i]
		if ft == nil {
			continue
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := castFolder(ft)
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			p, err := readFavrec(file)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.ThisFolder = p
			fav.FolderID++
			favFolder.Fid = fav.FolderID
		case FAVT_LINE:
			favLine := castLine(ft)
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			fav.LineID++
			favLine.Lid = fav.LineID
		}
	}

	return fav, nil
}

func castFolder(ft *FavType) *FavFolder {
	if ft.Fp != nil {
		favFolder, ok := ft.Fp.(*FavFolder)
		if !ok {
			return nil
		}
		return favFolder
	}
	favFolder := &FavFolder{}
	ft.Fp = favFolder
	return favFolder
}

func castBoard(ft *FavType) *FavBoard {
	if ft.Fp != nil {
		favBoard, ok := ft.Fp.(*FavBoard)
		if !ok {
			return nil
		}
		return favBoard
	}

	favBoard := &FavBoard{}
	ft.Fp = favBoard
	return favBoard
}

func castLine(ft *FavType) *FavLine {
	if ft.Fp != nil {
		favLine, ok := ft.Fp.(*FavLine)
		if !ok {
			return nil
		}
		return favLine
	}

	favLine := &FavLine{}
	ft.Fp = favLine
	return favLine
}

func isValidFavType(theType int8) bool {
	switch theType {
	case FAVT_BOARD:
	case FAVT_FOLDER:
	case FAVT_LINE:
		return true
	}
	return false
}

func getDataNumber(fav *Fav) int16 {
	return fav.NBoards + int16(fav.NLines) + int16(fav.NFolders)
}

func getTypeSize(theType int8) uintptr {
	switch theType {
	case FAVT_BOARD:
		return SIZE_OF_FAV_BOARD
	case FAVT_FOLDER:
		return SIZE_OF_FAV_FOLDER
	case FAVT_LINE:
		return SIZE_OF_FAV_LINE
	}
	return 0
}
