package fav

import (
	"encoding/binary"
	"os"
	"time"
	"unsafe"

	"github.com/PichuChen/go-bbs/path"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

const (
	_OFFSET_SAVE_MILLI_TS = 10000 //should not save the fav-board again within 10 sec.
)

type FavRaw struct {
	LoadTime int64
	//XXX fav is in-mem-cache.
	//    How can we notify the other same-user about the update of the fav？～
	//    kill the ssh/telnet/ws user？
	//    It should be the same user, the same user should be responsible for this behavior？～
	//	  Can we use load-time and save-time to update when the load-time > save-time? (require modification in C-part)
	NBoards  int16 /* number of the boards */
	NLines   int8  /* number of the lines */
	NFolders int8  /* number of the folders */
	LineID   int8  /* current max line id */
	FolderID int8  /* current max folder id */

	Favh []*FavType
}

const SIZE_OF_FAV = unsafe.Sizeof(FavRaw{})

func MTime(userID *[ptttype.IDLEN + 1]byte) (int64, error) {
	filename, err := path.SetHomeFile(userID, FAV)
	if err != nil {
		return 0, err
	}

	stat, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return 0, err
	}

	return stat.ModTime().UnixNano(), nil
}

//Load
//
//Load fav from file.
func Load(userID *[ptttype.IDLEN + 1]byte) (*FavRaw, error) {
	filename, err := path.SetHomeFile(userID, FAV)
	if err != nil {
		return nil, err
	}

	if !types.IsRegularFile(filename) {
		if true {
			fav, err := tryFav4Load(userID, filename)
			if err != nil {
				return nil, err
			}

			return fav, nil
		}
		result := &FavRaw{}
		result.LoadTime = types.GetCurrentMilliTS()
		return result, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// version
	version := int16(0)
	err = binary.Read(file, binary.LittleEndian, &version)
	if err != nil {
		return nil, err
	}
	// if version != FAV_VERSION {
	// }

	// favrec
	favrec, err := readFavrec(file)
	if err != nil {
		return nil, err
	}

	favrec.LoadTime = types.GetCurrentMilliTS()

	return favrec, nil
}

//Save
//
//save fav to file.
//XXX use rename to reduce the probability of race-condition.
func (fav *FavRaw) Save(userID *[ptttype.IDLEN + 1]byte) (*FavRaw, int64, error) {
	fav.cleanup()

	filename, err := path.SetHomeFile(userID, FAV)
	if err != nil {
		return nil, 0, err
	}

	// It's possible that the file does not exists.
	stat, err := os.Stat(filename)
	var mtime int64
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return nil, 0, err
	} else {
		mtime = stat.ModTime().UnixNano()
	}
	if mtime > (fav.LoadTime-_OFFSET_SAVE_MILLI_TS)*types.MILLI_TS_TO_NANO_TS {
		log.Errorf("%v fav.Save: too old: stat: %v loadTime: %v", time.Now(), mtime, (fav.LoadTime-_OFFSET_SAVE_MILLI_TS)*types.MILLI_TS_TO_NANO_TS)
		newFav, err := Load(userID)
		if err != nil {
			return nil, 0, err
		}
		return newFav, mtime, ErrOutdatedFav
	}

	postfix := types.GetRandom()
	tmpFilename, err := path.SetHomeFile(userID, FAV+".tmp."+postfix)
	if err != nil {
		return nil, 0, err
	}

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to create file")

	file, err := os.Create(tmpFilename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to write version")

	version := FAV_VERSION
	err = binary.Write(file, binary.LittleEndian, &version)
	if err != nil {
		return nil, 0, err
	}

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to write file")

	err = fav.writeFavrec(file)
	if err != nil {
		return nil, 0, err
	}

	log.WithFields(log.Fields{"tmpFilename": tmpFilename, "filename": filename}).Infof("to rename file")

	err = os.Rename(tmpFilename, filename)
	if err != nil {
		return nil, 0, err
	}

	newFav, err := Load(userID)
	if err != nil {
		return nil, 0, err
	}

	return newFav, mtime, err
}

func (fav *FavRaw) getDataNumber() int16 {
	return fav.NBoards + int16(fav.NLines) + int16(fav.NFolders)
}

func (fav *FavRaw) cleanup() {
	if !fav.isNeedRebuildFav() {
		return
	}

	fav.rebuildFav()
}

func (fav *FavRaw) isNeedRebuildFav() bool {
	if fav.Favh == nil {
		return false
	}
	for _, ft := range fav.Favh {
		if ft == nil {
			continue
		}

		if !ft.isValid() {
			return true
		}

		switch ft.TheType {
		case FAVT_BOARD:
		case FAVT_LINE:
		case FAVT_FOLDER:
			childFav := ft.castFolder().ThisFolder
			if childFav.isNeedRebuildFav() {
				return true
			}
		default:
			return true
		}
	}
	return false
}

/**
 * 清除 fp(dir) 中無效的 entry/dir。「無效」指的是沒有 FAVH_FAV flag，所以
 * 不包含不存在的看板。
 */
func (fav *FavRaw) rebuildFav() {
	fav.LineID = 0
	fav.FolderID = 0
	fav.NLines = 0
	fav.NFolders = 0
	fav.NBoards = 0

	if fav.Favh == nil {
		return
	}

	j := 0
	for i, ft := range fav.Favh {
		if !ft.isValid() {
			continue
		}

		switch ft.TheType {
		case FAVT_BOARD:
		case FAVT_LINE:
		case FAVT_FOLDER:
			childFav := ft.castFolder().ThisFolder
			childFav.rebuildFav()
		default:
			continue
		}

		fav.increase(ft)
		if i != j {
			ft.copyTo(fav.Favh[j])
		}
		j++
	}

	nFavh := fav.getDataNumber()
	// to be consistant with the data-tail.
	fav.Favh = fav.Favh[:nFavh]
}

func (fav *FavRaw) increase(ft *FavType) {
	switch ft.TheType {
	case FAVT_BOARD:
		fav.NBoards++
	case FAVT_LINE:
		fav.NLines++
		fav.LineID++
		ftLine := ft.castLine()
		ftLine.Lid = fav.LineID
	case FAVT_FOLDER:
		fav.NFolders++
		fav.FolderID++
		ftFolder := ft.castFolder()
		ftFolder.Fid = fav.FolderID
	}
}

func (fav *FavRaw) writeFavrec(file *os.File) error {
	if fav == nil {
		return nil
	}

	err := binary.Write(file, binary.LittleEndian, &fav.NBoards)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &fav.NLines)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &fav.NFolders)
	if err != nil {
		return err
	}
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
			err = binary.Write(file, binary.LittleEndian, &favFolder.Title)
			if err != nil {
				return err
			}
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return ErrInvalidFavBoard
			}
			err = types.BinWrite(file, favBoard, ft.TheType.GetTypeSize())
			if err != nil {
				return err
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return ErrInvalidFavLine
			}
			err = types.BinWrite(file, favLine, ft.TheType.GetTypeSize())
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
			err := favFolder.ThisFolder.writeFavrec(file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func tryFav4Load(userID *[ptttype.IDLEN + 1]byte, filename string) (*FavRaw, error) {
	oldFilename, err := path.SetHomeFile(userID, FAV4)
	if err != nil {
		return nil, err
	}
	if !types.IsRegularFile(oldFilename) {
		result := &FavRaw{}
		result.LoadTime = types.GetCurrentMilliTS()
		return result, nil
	}

	file, err := os.Open(oldFilename)
	if err != nil {
		return nil, err
	}

	fav, err := fav4ReadFavrec(file)
	if err != nil {
		return nil, err
	}
	_, _, _ = fav.Save(userID)

	bakFilename, err := path.SetHomeFile(userID, FAV+".bak")
	if err != nil {
		return nil, err
	}
	// XXX copy new fav-filename to bak in pttbbs
	err = types.CopyFile(filename, bakFilename)
	if err != nil {
		log.WithFields(log.Fields{"filename": filename, "bakFilename": bakFilename, "e": err}).Warn("unable to CopyFile")
	}

	return fav, nil
}

func fav4ReadFavrec(file *os.File) (*FavRaw, error) {
	fav := &FavRaw{}
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

	nFavh := fav.getDataNumber()
	fav.LineID = 0
	fav.FolderID = 0
	fav.Favh = make([]*FavType, nFavh)

	for i := int16(0); i < nFavh; i++ {
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
			err = types.BinRead(file, favFolder, ft.TheType.GetFav4TypeSize())
			if err != nil {
				return nil, ErrInvalidFav4Record
			}
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return nil, ErrInvalidFav4Record
			}
			err = types.BinRead(file, favBoard, ft.TheType.GetFav4TypeSize())
			if err != nil {
				return nil, ErrInvalidFav4Record
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFav4Record
			}
			err = types.BinRead(file, favLine, ft.TheType.GetFav4TypeSize())
			if err != nil {
				return nil, ErrInvalidFav4Record
			}
		}
	}

	for _, ft := range fav.Favh {
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

func readFavrec(file *os.File) (*FavRaw, error) {
	// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c
	fav := &FavRaw{}
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

	nFavh := fav.getDataNumber()
	fav.LineID = 0
	fav.FolderID = 0
	fav.Favh = make([]*FavType, nFavh)

	log.WithFields(log.Fields{"total": nFavh}).Debug("to for-loop")

	for i := int16(0); i < nFavh; i++ {
		ft := &FavType{}
		fav.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		log.WithFields(log.Fields{"type": ft.TheType, "e": err}).Debugf("(%v/%v) after read-type", i, nFavh)
		if err != nil {
			return nil, ErrInvalidFavType
		}
		if !ft.TheType.IsValidFavType() {
			return nil, ErrInvalidFavType
		}

		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		log.WithFields(log.Fields{"Attr": ft.Attr, "e": err}).Debugf("(%v/%v) after read-attr", i, nFavh)
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
			log.WithFields(log.Fields{"favFolder": favFolder, "e": err}).Debugf("(%v/%v) after read-fid", i, nFavh)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.Title)
			log.WithFields(log.Fields{"favFolder": favFolder, "e": err}).Debugf("(%v/%v) after read-title-big5", i, nFavh)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			log.WithFields(log.Fields{"favFolder": favFolder}).Debugf("(%v/%v) after read-folder", i, nFavh)
		case FAVT_BOARD:
			favBoard := ft.castBoard()
			if favBoard == nil {
				return nil, ErrInvalidFavBoard
			}
			err = types.BinRead(file, favBoard, ft.TheType.GetTypeSize())
			log.WithFields(log.Fields{"favBoard": favBoard, "size": SIZE_OF_FAV_BOARD}).Debugf("(%v/%v) after read-board", i, nFavh)
			if err != nil {
				return nil, ErrInvalidFavBoard
			}
		case FAVT_LINE:
			favLine := ft.castLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			err = types.BinRead(file, favLine, ft.TheType.GetTypeSize())
			if err != nil {
				return nil, ErrInvalidFavLine
			}
			log.WithFields(log.Fields{"favLine": favLine, "size": SIZE_OF_FAV_LINE}).Debugf("(%v/%v) after read-line", i, nFavh)
		}
	}

	for i := int16(0); i < nFavh; i++ {
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
