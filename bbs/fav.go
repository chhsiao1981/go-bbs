package bbs

import (
	"encoding/binary"
	"os"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

const (
	FAV_VERSION int16 = 3363
)

type FAVT int8

const (
	FAVT_BOARD  FAVT = 1
	FAVT_FOLDER FAVT = 2
	FAVT_LINE   FAVT = 3
)

func (t FAVT) String() string {
	switch t {
	case FAVT_BOARD:
		return "Board"
	case FAVT_FOLDER:
		return "Folder"
	case FAVT_LINE:
		return "Line"
	default:
		return "unknown"
	}
}

type FAVH int8

const (
	FAVH_FAV     FAVH = 1
	FAVH_TAG     FAVH = 2
	FAVH_UNREAD  FAVH = 4
	FAVH_ADM_TAG FAVH = 8
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
	TheType FAVT
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
	Attr      int8
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

type Fav4Folder struct {
	Fid        int8
	TitleBig5  [BTLEN + 1]byte
	ThisFolder int32
}

const SIZE_OF_FAV4_FOLDER = unsafe.Sizeof(Fav4Folder{})

type Fav4Board struct {
	Bid       int32
	LastVisit int32
	Attr      int8
}

const SIZE_OF_FAV4_BOARD = unsafe.Sizeof(Fav4Board{})

func FavSave(fav *Fav, userID string) error {
	filename := setuserfile(userID, FAV)
	postfix := getRandom().String()
	tmpFilename := setuserfile(userID, FAV+".tmp."+postfix)

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to create file")

	file, err := os.Create(tmpFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to write version")

	version := FAV_VERSION
	err = binary.Write(file, binary.LittleEndian, &version)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to write file")

	err = writeFavrec(file, fav)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to rename file")

	err = os.Rename(tmpFilename, filename)
	if err != nil {
		return err
	}

	return nil
}

func writeFavrec(file *os.File, fav *Fav) error {
	if fav == nil {
		return nil
	}

	binary.Write(file, binary.LittleEndian, &fav.NBoards)
	binary.Write(file, binary.LittleEndian, &fav.NLines)
	binary.Write(file, binary.LittleEndian, &fav.NFolders)
	total := fav.getDataNumber()
	for i := int16(0); i < total; i++ {
		ft := fav.Favh[i]
		err := binary.Write(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return err
		}
		err = binary.Write(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return err
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := ft.castFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err = binary.Write(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return err
			}
			err = binary.Write(file, binary.LittleEndian, &favFolder.TitleBig5)
			if err != nil {
				return err
			}
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return ErrInvalidFavBoard
			}
			err = binWrite(file, favBoard, getTypeSize(ft.TheType))
			if err != nil {
				return err
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return ErrInvalidFavLine
			}
			err = binWrite(file, favLine, getTypeSize(ft.TheType))
			if err != nil {
				return err
			}
		}
	}

	for i := int16(0); i < total; i++ {
		ft := fav.Favh[i]
		if ft == nil {
			continue
		}
		if ft.TheType == FAVT_FOLDER {
			favFolder := ft.castFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err := writeFavrec(file, favFolder.ThisFolder)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func FavLoad(userID string) (*Fav, error) {
	filename := setuserfile(userID, FAV)

	if !isRegularFile(filename) {
		if true {
			fav, err := tryFav4Load(userID, filename)
			if err != nil {
				return nil, err
			}

			return fav, nil
		}
		return &Fav{}, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// version
	version := int16(0)
	binary.Read(file, binary.LittleEndian, &version)
	// if version != FAV_VERSION {
	// }

	// favrec
	favrec, err := readFavrec(file)
	if err != nil {
		return nil, err
	}

	return favrec, nil
}

func tryFav4Load(userID string, filename string) (*Fav, error) {
	oldFilename := setuserfile(userID, FAV4)
	if !isRegularFile(oldFilename) {
		return &Fav{}, nil
	}

	file, err := os.Open(oldFilename)
	if err != nil {
		return nil, err
	}

	fav, err := fav4ReadFavrec(file)
	if err != nil {
		return nil, err
	}
	_ = FavSave(fav, userID)

	bakFilename := setuserfile(userID, FAV+".bak")
	// XXX copy new fav-filename to bak in pttbbs
	_, err = CopyFile(filename, bakFilename)
	if err != nil {
		log.WithFields(log.Fields{"filename": filename, "bakFilename": bakFilename, "e": err}).Warn("unable to CopyFile")
	}

	return fav, nil
}

func fav4ReadFavrec(file *os.File) (*Fav, error) {
	fav := &Fav{}
	err := binary.Read(file, binary.LittleEndian, &fav.NBoards)
	if err != nil {
		return nil, ErrInvalidFav4Record
	}

	err = binary.Read(file, binary.LittleEndian, &fav.NLines)
	if err != nil {
		return nil, ErrInvalidFav4Record
	}

	err = binary.Read(file, binary.LittleEndian, &fav.NFolders)
	if err != nil {
		return nil, ErrInvalidFav4Record
	}

	fav.DataTail = fav.getDataNumber()
	fav.NAllocs = fav.DataTail + FAV_PRE_ALLOC
	fav.LineID = 0
	fav.FolderID = 0
	fav.Favh = make([]*FavType, fav.NAllocs)

	for i := int16(0); i < fav.DataTail; i++ {
		ft := &FavType{}
		fav.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return nil, ErrInvalidFav4Record
		}
		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return nil, ErrInvalidFav4Record
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := ft.castFolder()
			if favFolder == nil {
				return nil, ErrInvalidFav4Record
			}
			err = binRead(file, favFolder, getFav4TypeSize(ft.TheType))
			if err != nil {
				return nil, ErrInvalidFav4Record
			}
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return nil, ErrInvalidFav4Record
			}
			err = binRead(file, favBoard, getFav4TypeSize(ft.TheType))
			if err != nil {
				return nil, ErrInvalidFav4Record
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFav4Record
			}
			err = binRead(file, favLine, getFav4TypeSize(ft.TheType))
			if err != nil {
				return nil, ErrInvalidFav4Record
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
			favFolder := ft.castFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			p, err := fav4ReadFavrec(file)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.ThisFolder = p
			fav.FolderID++
			favFolder.Fid = fav.FolderID
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			fav.LineID++
			favLine.Lid = fav.LineID
		}
	}

	return fav, nil
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

	fav.DataTail = fav.getDataNumber()
	fav.NAllocs = fav.DataTail + FAV_PRE_ALLOC
	fav.LineID = 0
	fav.FolderID = 0
	fav.Favh = make([]*FavType, fav.NAllocs)

	log.WithFields(log.Fields{"total": fav.DataTail}).Debug("to for-loop")

	for i := int16(0); i < fav.DataTail; i++ {
		ft := &FavType{}
		fav.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		log.WithFields(log.Fields{"type": ft.TheType, "e": err}).Debugf("(%v/%v) after read-type", i, fav.DataTail)
		if err != nil {
			return nil, ErrInvalidFavType
		}
		if !isValidFavType(ft.TheType) {
			return nil, ErrInvalidFavType
		}

		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		log.WithFields(log.Fields{"Attr": ft.Attr, "e": err}).Debugf("(%v/%v) after read-attr", i, fav.DataTail)
		if err != nil {
			return nil, ErrInvalidFavType
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := ft.castFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.Fid)
			log.WithFields(log.Fields{"favFolder": favFolder, "e": err}).Debugf("(%v/%v) after read-fid", i, fav.DataTail)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.TitleBig5)
			log.WithFields(log.Fields{"favFolder": favFolder, "e": err}).Debugf("(%v/%v) after read-title-big5", i, fav.DataTail)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			log.WithFields(log.Fields{"favFolder": favFolder}).Debugf("(%v/%v) after read-folder", i, fav.DataTail)
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return nil, ErrInvalidFavBoard
			}
			err = binRead(file, favBoard, getTypeSize(ft.TheType))
			log.WithFields(log.Fields{"favBoard": favBoard, "size": SIZE_OF_FAV_BOARD}).Debugf("(%v/%v) after read-board", i, fav.DataTail)
			if err != nil {
				return nil, ErrInvalidFavBoard
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			err = binRead(file, favLine, getTypeSize(ft.TheType))
			if err != nil {
				return nil, ErrInvalidFavLine
			}
			log.WithFields(log.Fields{"favLine": favLine, "size": SIZE_OF_FAV_LINE}).Debugf("(%v/%v) after read-line", i, fav.DataTail)
		}
	}

	for i := int16(0); i < fav.DataTail; i++ {
		ft := fav.Favh[i]
		if ft == nil {
			continue
		}

		switch ft.TheType {
		case FAVT_FOLDER:
			favFolder := ft.castFolder()
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
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			fav.LineID++
			favLine.Lid = fav.LineID
		}
	}

	return fav, nil
}

func (ft *FavType) castFolder() *FavFolder {
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

func (ft *FavType) castBoard() *FavBoard {
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

func (ft *FavType) castLine() *FavLine {
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

func isValidFavType(theType FAVT) bool {
	switch theType {
	case FAVT_BOARD:
		return true
	case FAVT_FOLDER:
		return true
	case FAVT_LINE:
		return true
	}
	return false
}

func (fav *Fav) getDataNumber() int16 {
	return fav.NBoards + int16(fav.NLines) + int16(fav.NFolders)
}

func getTypeSize(theType FAVT) uintptr {
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

func getFav4TypeSize(theType FAVT) uintptr {
	switch theType {
	case FAVT_BOARD:
		return SIZE_OF_FAV4_BOARD
	case FAVT_FOLDER:
		return SIZE_OF_FAV_FOLDER
	case FAVT_LINE:
		return SIZE_OF_FAV_LINE
	}
	return 0
}
