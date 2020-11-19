package cache

import (
	"bytes"
	"unsafe"

	"github.com/PichuChen/go-bbs/cmsys"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

//AddToUHash
func AddToUHash(uidInCache int32, userID *[ptttype.IDLEN + 1]byte) error {
	h := cmsys.StringHashWithHashBits(userID[:])

	// line: 166
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.Userid)+ptttype.USER_ID_SZ*uintptr(uidInCache),
		ptttype.USER_ID_SZ,
		unsafe.Pointer(userID),
	)

	// init vars
	p := h
	val := int32(0)
	pval := &val
	valptr := unsafe.Pointer(pval)
	offset := unsafe.Offsetof(Shm.Raw.HashHead)

	// line: 168
	Shm.ReadAt(
		offset+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		valptr,
	)

	times := 0
	offsetNextInHash := unsafe.Offsetof(Shm.Raw.NextInHash)
	for ; times < ptttype.MAX_USERS && val != -1; times++ {
		offset = offsetNextInHash
		p = uint32(val)
		Shm.ReadAt(
			offset+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			valptr,
		)
	}
	if times >= ptttype.MAX_USERS {
		log.Errorf("Unable to add-to-uhash! uid-in-cache: %v userID: %v", uidInCache, string(userID[:]))
		return ErrAddToUHash
	}

	// set current ptr
	*pval = uidInCache
	Shm.WriteAt(
		offset+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		valptr,
	)

	// set next as -1
	*pval = -1
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.NextInHash)+types.INT32_SZ*uintptr(uidInCache),
		types.INT32_SZ,
		valptr,
	)

	return nil
}

//RemoveFromUHash
func RemoveFromUHash(uidInCache int32) error {
	userID := &[ptttype.IDLEN + 1]byte{}

	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.Userid)+ptttype.USER_ID_SZ*uintptr(uidInCache),
		ptttype.USER_ID_SZ,
		unsafe.Pointer(userID),
	)

	h := cmsys.StringHashWithHashBits(userID[:])

	// line: 191
	p := h
	val := int32(0)
	pval := &val
	valptr := unsafe.Pointer(pval)
	offset := unsafe.Offsetof(Shm.Raw.HashHead)
	Shm.ReadAt(
		offset+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		valptr,
	)

	// line: 194
	times := 0
	for ; times < ptttype.MAX_USERS && val != -1 && val != uidInCache; times++ {
		p = uint32(val)
		offset = unsafe.Offsetof(Shm.Raw.NextInHash)
		Shm.ReadAt(
			offset+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			valptr,
		)
	}
	if times >= ptttype.MAX_USERS {
		log.Errorf("Unable to remove-from-uhash! uid-in-cache: %v userID: %v", uidInCache, string(userID[:]))
		return ErrRemoveFromUHash
	}

	if val == uidInCache {
		nextNum := int32(0)
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.NextInHash)+types.INT32_SZ*uintptr(uidInCache),
			types.INT32_SZ,
			unsafe.Pointer(&nextNum),
		)

		*pval = nextNum
		Shm.WriteAt(
			offset+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			valptr,
		)
	}
	return nil
}

//SearchUser
//Params:
//	userID: querying user-id.
//	isReturn: is return the user-id in the shm.
//
//Return:
//	int: uid.
//	string: the userID in shm.
//	error: err.
func SearchUser(userID string, isReturn bool) (uid int32, rightID string, err error) {
	if len(userID) == 0 {
		return 0, "", nil
	}
	return doSearchUser(userID, isReturn)
}

//doSearchUser
//Params:
//	userID
//	isReturn
//
//Return:
//	int32: uid
//	string: the userID in shm.
//	error: err.
func doSearchUser(userID string, isReturn bool) (uid int32, rightID string, err error) {
	userIDBytes := &[ptttype.IDLEN + 1]byte{}
	copy(userIDBytes[:], []byte(userID))

	var rightIDBytes *[ptttype.IDLEN + 1]byte = nil
	if isReturn {
		rightIDBytes = &[ptttype.IDLEN + 1]byte{}
	}

	uid, err = doSearchUserRaw(userIDBytes, rightIDBytes)
	if err != nil {
		return 0, "", err
	}

	rightID = ""
	if isReturn {
		rightID = types.CstrToString(rightIDBytes[:])
	}

	return uid, rightID, nil
}

//SearchUser
//Params:
//	userID: querying user-id.
//	isReturn: is return the user-id in the shm.
//
//Return:
//	int: uid.
//	string: the userID in shm.
//	error: err.
func SearchUserRaw(userID *[ptttype.IDLEN + 1]byte, rightID *[ptttype.IDLEN + 1]byte) (uid int32, err error) {
	if userID[0] == 0 {
		return 0, nil
	}
	return doSearchUserRaw(userID, rightID)
}

func doSearchUserRaw(userID *[ptttype.IDLEN + 1]byte, rightID *[ptttype.IDLEN + 1]byte) (int32, error) {
	// XXX we should have 0 as non-exists.
	//     currently the reason why it's ok is because the probability of collision on 0 is low.

	_ = StatInc(ptttype.STAT_SEARCHUSER)
	h := cmsys.StringHashWithHashBits(userID[:])

	//p = SHM->hash_head[h]  //line: 219
	p := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.HashHead)+types.INT32_SZ*uintptr(h),
		types.INT32_SZ,
		unsafe.Pointer(&p),
	)
	log.Debugf("doSearchUserRaw: after Shm.ReadAt: userID: %v h: %v p: %v", string(userID[:]), h, p)

	shmUserID := [ptttype.IDLEN + 1]byte{}

	for times := 0; times < ptttype.MAX_USERS && p != -1 && p < ptttype.MAX_USERS; times++ {
		//if (strcasecmp(SHM->userid[p], userid) == 0)  //line: 222
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.Userid)+ptttype.USER_ID_SZ*uintptr(p),
			ptttype.USER_ID_SZ,
			unsafe.Pointer(&shmUserID),
		)
		if bytes.Equal(bytes.ToUpper(userID[:]), bytes.ToUpper(shmUserID[:])) {
			if userID[0] != 0 && rightID != nil {
				copy(rightID[:], shmUserID[:])
				log.Infof("doSearchUserRaw: after copy: rightID: %v shmUserID: %v", string(rightID[:]), string(shmUserID[:]))
			}
			return p + 1, nil
		}
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.NextInHash)+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			unsafe.Pointer(&p),
		)
	}

	return 0, nil
}

//GetUserID
//
//XXX uid = uid-in-cache + 1
func GetUserID(uid int32) (*[ptttype.IDLEN + 1]byte, error) {
	uid--
	if uid < 0 || uid >= ptttype.MAX_USERS {
		return nil, ErrInvalidUID
	}

	userID := &[ptttype.IDLEN + 1]byte{}
	log.Infof("GetUserID: to ReadAt: uid: %v", uid)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.Userid)+ptttype.USER_ID_SZ*uintptr(uid),
		ptttype.USER_ID_SZ,
		unsafe.Pointer(userID),
	)

	return userID, nil
}

//SetUserID
//
//XXX uid = uid-in-cache + 1
func SetUserID(uid int32, userID *[ptttype.IDLEN + 1]byte) (err error) {
	if uid <= 0 || uid > ptttype.MAX_USERS {
		return ErrInvalidUID
	}

	errRemove := RemoveFromUHash(uid - 1)
	errAdd := AddToUHash(uid-1, userID)
	if errRemove != nil {
		return errRemove
	}
	if errAdd != nil {
		return errAdd
	}

	return nil
}

func SearchUListUserID(userID *[ptttype.IDLEN + 1]byte) *ptttype.UserInfoRaw {
	// start and end
	start := int32(0)

	end := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.UTMPNumber),
		types.INT32_SZ,
		unsafe.Pointer(&end),
	)
	end--
	if end < 0 {
		return nil
	}

	// current-sorted
	currentSorted := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Raw.CurrSorted),
		types.INT32_SZ,
		unsafe.Pointer(&currentSorted),
	)

	// search
	ulist_i := int32(0)
	pulist_i := &ulist_i
	ulist_iptr := unsafe.Pointer(pulist_i)
	isDiff := 0
	const offsetUInfoUserID = unsafe.Offsetof(Shm.Raw.UInfo[0].UserID)
	const sizeOfSorted = unsafe.Sizeof(Shm.Raw.Sorted[0])
	const sizeOfSorted2 = unsafe.Sizeof(Shm.Raw.Sorted[0][0])
	uInfo := &ptttype.UserInfoRaw{}
	for i := (start + end) / 2; ; i = (start + end) / 2 {
		// get ulist_i
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Raw.Sorted)+sizeOfSorted*uintptr(currentSorted)+sizeOfSorted2*0+types.INT32_SZ*uintptr(i),
			types.INT32_SZ,
			ulist_iptr,
		)

		// do cmp.
		isDiff = Shm.Cmp(
			unsafe.Offsetof(Shm.Raw.UInfo)+ptttype.USER_INFO_RAW_SZ*uintptr(ulist_i)+offsetUInfoUserID,
			ptttype.USER_ID_SZ,
			unsafe.Pointer(userID),
		)
		if isDiff == 0 {
			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.UInfo)+ptttype.USER_INFO_RAW_SZ*uintptr(ulist_i),
				ptttype.USER_INFO_RAW_SZ,
				unsafe.Pointer(uInfo),
			)

			return uInfo
		}

		// determine start / end
		if end == start {
			break
		} else if i == start {
			start = end
		} else if isDiff > 0 {
			start = i
		} else {
			end = i
		}
	}

	return nil
}
