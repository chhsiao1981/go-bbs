package cache

import (
	"bufio"
	"bytes"
	"os"
	"unsafe"

	"github.com/PichuChen/go-bbs/path"
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
		theList := bytes.Split(line, []byte{' '})

		eachUserID := &ptttype.UserID_t{}
		copy(eachUserID[:], theList[0][:])

		if bytes.EqualFold(eachUserID[:], ptttype.USER_ID_GUEST[:]) {
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
