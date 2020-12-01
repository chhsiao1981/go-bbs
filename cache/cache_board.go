package cache

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"time"
	"unsafe"

	"github.com/PichuChen/go-bbs/cmbbs/path"
	"github.com/PichuChen/go-bbs/cmsys"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func GetBCache(bidInCache int32) (board *ptttype.BoardHeaderRaw, err error) {
	board = &ptttype.BoardHeaderRaw{}

	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.BCache)+ptttype.BOARD_HEADER_RAW_SZ*uintptr(bidInCache),
		ptttype.BOARD_HEADER_RAW_SZ,
		unsafe.Pointer(board),
	)
	return board, nil
}

func IsHiddenBoardFriend(bidInCache int32, uidInCache int32) bool {
	//hbfl time
	var loadTime types.Time4
	pLoadTime := &loadTime

	const Hbfl0Size = unsafe.Sizeof(Shm.Raw.Hbfl[0])
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.Hbfl)+Hbfl0Size*uintptr(bidInCache),
		types.TIME4_SZ,
		unsafe.Pointer(pLoadTime),
	)

	// XXX use nowTS to replace loginStartTime.
	//     HBFLexpire is set as 5-days. nowTS should be ok.
	nowTS := types.NowTS()
	if loadTime < nowTS-types.Time4(ptttype.HBFLexpire) {
		HbflReload(bidInCache)
	}

	uid := uidInCache + 1

	var friendID int32
	pFriendID := &friendID
	friendIDptr := unsafe.Pointer(pFriendID)
	for i := uintptr(1); i <= ptttype.MAX_FRIEND; i++ {
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.Hbfl)+Hbfl0Size*uintptr(bidInCache)+types.INT32_SZ*i,
			types.INT32_SZ,
			friendIDptr,
		)
		if friendID == 0 {
			break
		}

		if friendID == uid {
			return true
		}
	}

	return false
}

func HbflReload(bidInCache int32) {
	brdname := &ptttype.BoardID_t{}

	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.BCache)+ptttype.BOARD_HEADER_RAW_SZ*uintptr(bidInCache)+ptttype.BOARD_HEADER_BOARD_NAME_OFFSET,
		ptttype.BOARD_ID_SZ,
		unsafe.Pointer(brdname),
	)

	filename, err := path.SetBFile(brdname, ptttype.FN_VISIBLE)
	if err != nil {
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	hbfl := [ptttype.MAX_FRIEND + 1]int32{}
	const hbflsz = unsafe.Sizeof(hbfl)

	scanner := bufio.NewScanner(file)
	var line []byte
	var uid int32
	// num++ is in the end of the for.
	for num := 1; scanner.Scan() && num <= ptttype.MAX_FRIEND; {

		line = scanner.Bytes()
		theList := bytes.Split(line, []byte{' '}) //The \x00 is taken care of by scanner.

		eachUserID := &ptttype.UserID_t{}
		copy(eachUserID[:], theList[0][:])

		if types.Cstrcasecmp(eachUserID[:], ptttype.USER_ID_GUEST[:]) == 0 {
			continue
		}

		uid, err = SearchUserRaw(eachUserID, nil)
		if err != nil {
			continue
		}
		if uid == 0 {
			continue
		}

		hbfl[num] = uid

		num++ // num++ is in the end of the for.
	}

	hbfl[0] = int32(types.NowTS())

	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.Hbfl)+hbflsz*uintptr(bidInCache),
		hbflsz,
		unsafe.Pointer(&hbfl),
	)
}

//NumBoards
//
//https://github.com/ptt/pttbbs/blob/master/common/bbs/cache.c#L512
func NumBoards() int32 {
	var nboards int32
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.BNumber),
		types.INT32_SZ,
		unsafe.Pointer(&nboards),
	)

	return nboards
}

//Reload BCache
//
//https://github.com/ptt/pttbbs/blob/master/common/bbs/cache.c#L458
func ReloadBCache() {
	var busystate int32
	for i := 0; i < 10; i++ { //Is it ok that we don't use mutex or semaphore here?
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.BBusyState),
			types.INT32_SZ,
			unsafe.Pointer(&busystate),
		)
		if busystate == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	// should we check that the busystate is still != 0 and return?

	busystate = 1
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BBusyState),
		types.INT32_SZ,
		unsafe.Pointer(&busystate),
	)

	theBytes, err := reloadBCacheReadFile()
	if err != nil {
		return
	}

	const bcachesz = unsafe.Sizeof(Shm.Raw.BCache)
	var theSize = bcachesz
	lenTheBytes := uintptr(len(theBytes))
	if lenTheBytes < theSize {
		theSize = lenTheBytes
	}
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BCache),
		theSize,
		unsafe.Pointer(&theBytes),
	)

	bnumber := int32(theSize / ptttype.BOARD_HEADER_RAW_SZ)
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BNumber),
		types.INT32_SZ,
		unsafe.Pointer(&bnumber),
	)

	Shm.Memset(
		unsafe.Offsetof(Shm.Raw.LastPostTime),
		byte(0),
		uintptr(ptttype.MAX_BOARD)*types.TIME4_SZ,
	)

	Shm.Memset(
		unsafe.Offsetof(Shm.Raw.Total),
		byte(0),
		uintptr(ptttype.MAX_BOARD)*types.INT32_SZ,
	)

	Shm.InnerSetInt32(
		unsafe.Offsetof(Shm.Raw.BTouchTime),
		unsafe.Offsetof(Shm.Raw.BUptime),
	)

	busystate = 0
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BBusyState),
		types.INT32_SZ,
		unsafe.Pointer(&busystate),
	)

	sortBCache()

	go reloadCacheLoadBottom()
}

//sortBCache
//XXX TODO: implement
func sortBCache() {
	var busystate int32
	pbusystate := &busystate
	pbusystateptr := unsafe.Pointer(pbusystate)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.BBusyState),
		types.INT32_SZ,
		pbusystateptr,
	)

	if busystate != 0 {
		time.Sleep(1 * time.Second)
		return
	}

	*pbusystate = 1
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BBusyState),
		types.INT32_SZ,
		pbusystateptr,
	)
	defer func() {
		*pbusystate = 0
		Shm.WriteAt(
			unsafe.Offsetof(Shm.Raw.BBusyState),
			types.INT32_SZ,
			pbusystateptr,
		)
	}()

	Shm.QsortCmpBoardName()
	Shm.QsortCmpBoardClass()

	// for-loop cleaning first-child
}

func reloadCacheLoadBottom() {
	boardName := &ptttype.BoardID_t{}
	for i := uintptr(0); i < ptttype.MAX_BOARD; i++ {
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.BCache)+ptttype.BOARD_HEADER_RAW_SZ*i+ptttype.BOARD_HEADER_BOARD_NAME_OFFSET,
			ptttype.BOARD_HEADER_RAW_SZ,
			unsafe.Pointer(boardName),
		)

		if boardName[0] == 0 {
			continue
		}

		filename, err := path.SetBFile(boardName, ptttype.FN_DIR_BOTTOM)
		if err != nil {
			continue
		}

		n := cmsys.GetNumRecords(filename, ptttype.FILE_HEADER_RAW_SZ)
		if n > 5 {
			n = 5
		}

		var n_uint8 = uint8(n)
		Shm.WriteAt(
			unsafe.Offsetof(Shm.Raw.NBottom)+i,
			1,
			unsafe.Pointer(&n_uint8),
		)
	}
}

func reloadBCacheReadFile() ([]byte, error) {
	file, err := os.Open(ptttype.FN_BOARD)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	theBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return theBytes, nil
}
