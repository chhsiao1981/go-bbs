package cache

import (
    "runtime/debug"
    "unsafe"

    "github.com/PichuChen/go-bbs/ptttype"
    "github.com/PichuChen/go-bbs/shm"
    "github.com/PichuChen/go-bbs/types"
    log "github.com/sirupsen/logrus"
)

type SHM struct {
    Shmid   int
    IsNew   bool
    Shmaddr unsafe.Pointer

    Raw SHMRaw //dummy variable to get the offset and size of the shm-fields.
}

//NewSHM
//
//This is to init SHM with Version and Size checked.
func NewSHM(key types.Key_t, isUseHugeTlb bool, isCreate bool) error {
    if Shm != nil {
        return ErrShmAlreadyInit
    }

    shmid := int(0)
    var shmaddr unsafe.Pointer
    isNew := false
    var err error

    // pttstruct.h line: 616
    SHMSIZE := types.Size_t(SHM_RAW_SZ)
    if ptttype.SHMALIGNEDSIZE != 0 {
        SHMSIZE = types.Size_t((int(SHM_RAW_SZ)/(ptttype.SHMALIGNEDSIZE) + 1) * ptttype.SHMALIGNEDSIZE)
    }

    log.Infof("NewSHM: SHMSIZE: %v SHM_RAW_SZ: %v SHMALIGNEDSIZE: %v", SHMSIZE, SHM_RAW_SZ, ptttype.SHMALIGNEDSIZE)

    log.Infof("NewSHM: SHMRaw.Version: (%v/%v)", unsafe.Offsetof(Shm.Raw.Version), unsafe.Sizeof(Shm.Raw.Version))

    log.Infof("NewSHM: SHMRaw.Size: (%v/%v)", unsafe.Offsetof(Shm.Raw.Size), unsafe.Sizeof(Shm.Raw.Size))

    log.Infof("NewSHM: SHMRaw.Userid: (%v/%v)", unsafe.Offsetof(Shm.Raw.Userid), unsafe.Sizeof(Shm.Raw.Userid))

    log.Infof("NewSHM: SHMRaw.Gap1: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap1), unsafe.Sizeof(Shm.Raw.Gap1))

    log.Infof("NewSHM: SHMRaw.NextInHash: (%v/%v)", unsafe.Offsetof(Shm.Raw.NextInHash), unsafe.Sizeof(Shm.Raw.NextInHash))

    log.Infof("NewSHM: SHMRaw.Gap2: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap2), unsafe.Sizeof(Shm.Raw.Gap2))

    log.Infof("NewSHM: SHMRaw.Money: (%v/%v)", unsafe.Offsetof(Shm.Raw.Money), unsafe.Sizeof(Shm.Raw.Money))

    log.Infof("NewSHM: SHMRaw.Gap3: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap3), unsafe.Sizeof(Shm.Raw.Gap3))

    log.Infof("NewSHM: SHMRaw.CooldownTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.CooldownTime), unsafe.Sizeof(Shm.Raw.CooldownTime))

    log.Infof("NewSHM: SHMRaw.Gap4: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap4), unsafe.Sizeof(Shm.Raw.Gap4))

    log.Infof("NewSHM: SHMRaw.HashHead: (%v/%v)", unsafe.Offsetof(Shm.Raw.HashHead), unsafe.Sizeof(Shm.Raw.HashHead))

    log.Infof("NewSHM: SHMRaw.Gap5: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap5), unsafe.Sizeof(Shm.Raw.Gap5))

    log.Infof("NewSHM: SHMRaw.Number: (%v/%v)", unsafe.Offsetof(Shm.Raw.Number), unsafe.Sizeof(Shm.Raw.Number))

    log.Infof("NewSHM: SHMRaw.Loaded: (%v/%v)", unsafe.Offsetof(Shm.Raw.Loaded), unsafe.Sizeof(Shm.Raw.Loaded))

    log.Infof("NewSHM: SHMRaw.UInfo: (%v/%v)", unsafe.Offsetof(Shm.Raw.UInfo), unsafe.Sizeof(Shm.Raw.UInfo))

    log.Infof("NewSHM: SHMRaw.Gap6: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap6), unsafe.Sizeof(Shm.Raw.Gap6))

    log.Infof("NewSHM: SHMRaw.Sorted: (%v/%v)", unsafe.Offsetof(Shm.Raw.Sorted), unsafe.Sizeof(Shm.Raw.Sorted))

    log.Infof("NewSHM: SHMRaw.Gap7: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap7), unsafe.Sizeof(Shm.Raw.Gap7))

    log.Infof("NewSHM: SHMRaw.CurrSorted: (%v/%v)", unsafe.Offsetof(Shm.Raw.CurrSorted), unsafe.Sizeof(Shm.Raw.CurrSorted))

    log.Infof("NewSHM: SHMRaw.UTMPUptime: (%v/%v)", unsafe.Offsetof(Shm.Raw.UTMPUptime), unsafe.Sizeof(Shm.Raw.UTMPUptime))

    log.Infof("NewSHM: SHMRaw.UTMPNumber: (%v/%v)", unsafe.Offsetof(Shm.Raw.UTMPNumber), unsafe.Sizeof(Shm.Raw.UTMPNumber))

    log.Infof("NewSHM: SHMRaw.UTMPNeedSort: (%v/%v)", unsafe.Offsetof(Shm.Raw.UTMPNeedSort), unsafe.Sizeof(Shm.Raw.UTMPNeedSort))

    log.Infof("NewSHM: SHMRaw.UTMPBusyState: (%v/%v)", unsafe.Offsetof(Shm.Raw.UTMPBusyState), unsafe.Sizeof(Shm.Raw.UTMPBusyState))

    log.Infof("NewSHM: SHMRaw.Gap8: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap8), unsafe.Sizeof(Shm.Raw.Gap8))

    log.Infof("NewSHM: SHMRaw.BMCache: (%v/%v)", unsafe.Offsetof(Shm.Raw.BMCache), unsafe.Sizeof(Shm.Raw.BMCache))

    log.Infof("NewSHM: SHMRaw.Gap9: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap9), unsafe.Sizeof(Shm.Raw.Gap9))

    log.Infof("NewSHM: SHMRaw.BCache: (%v/%v)", unsafe.Offsetof(Shm.Raw.BCache), unsafe.Sizeof(Shm.Raw.BCache))

    log.Infof("NewSHM: SHMRaw.Gap10: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap10), unsafe.Sizeof(Shm.Raw.Gap10))

    log.Infof("NewSHM: SHMRaw.BSorted: (%v/%v)", unsafe.Offsetof(Shm.Raw.BSorted), unsafe.Sizeof(Shm.Raw.BSorted))

    log.Infof("NewSHM: SHMRaw.Gap11: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap11), unsafe.Sizeof(Shm.Raw.Gap11))

    log.Infof("NewSHM: SHMRaw.NHOTs: (%v/%v)", unsafe.Offsetof(Shm.Raw.NHOTs), unsafe.Sizeof(Shm.Raw.NHOTs))

    log.Infof("NewSHM: SHMRaw.HBcache: (%v/%v)", unsafe.Offsetof(Shm.Raw.HBcache), unsafe.Sizeof(Shm.Raw.HBcache))

    log.Infof("NewSHM: SHMRaw.Gap12: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap12), unsafe.Sizeof(Shm.Raw.Gap12))

    log.Infof("NewSHM: SHMRaw.BusyStateB: (%v/%v)", unsafe.Offsetof(Shm.Raw.BusyStateB), unsafe.Sizeof(Shm.Raw.BusyStateB))

    log.Infof("NewSHM: SHMRaw.Gap13: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap13), unsafe.Sizeof(Shm.Raw.Gap13))

    log.Infof("NewSHM: SHMRaw.Total: (%v/%v)", unsafe.Offsetof(Shm.Raw.Total), unsafe.Sizeof(Shm.Raw.Total))

    log.Infof("NewSHM: SHMRaw.Gap14: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap14), unsafe.Sizeof(Shm.Raw.Gap14))

    log.Infof("NewSHM: SHMRaw.NBottom: (%v/%v)", unsafe.Offsetof(Shm.Raw.NBottom), unsafe.Sizeof(Shm.Raw.NBottom))

    log.Infof("NewSHM: SHMRaw.Gap15: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap15), unsafe.Sizeof(Shm.Raw.Gap15))

    log.Infof("NewSHM: SHMRaw.Hbfl: (%v/%v)", unsafe.Offsetof(Shm.Raw.Hbfl), unsafe.Sizeof(Shm.Raw.Hbfl))

    log.Infof("NewSHM: SHMRaw.Gap16: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap16), unsafe.Sizeof(Shm.Raw.Gap16))

    log.Infof("NewSHM: SHMRaw.LastPostTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.LastPostTime), unsafe.Sizeof(Shm.Raw.LastPostTime))

    log.Infof("NewSHM: SHMRaw.Gap17: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap17), unsafe.Sizeof(Shm.Raw.Gap17))

    log.Infof("NewSHM: SHMRaw.BUptime: (%v/%v)", unsafe.Offsetof(Shm.Raw.BUptime), unsafe.Sizeof(Shm.Raw.BUptime))

    log.Infof("NewSHM: SHMRaw.BTouchTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.BTouchTime), unsafe.Sizeof(Shm.Raw.BTouchTime))

    log.Infof("NewSHM: SHMRaw.Version: (%v/%v)", unsafe.Offsetof(Shm.Raw.Version), unsafe.Sizeof(Shm.Raw.Version))

    log.Infof("NewSHM: SHMRaw.BNumber: (%v/%v)", unsafe.Offsetof(Shm.Raw.BNumber), unsafe.Sizeof(Shm.Raw.BNumber))

    log.Infof("NewSHM: SHMRaw.BBusyState: (%v/%v)", unsafe.Offsetof(Shm.Raw.BBusyState), unsafe.Sizeof(Shm.Raw.BBusyState))

    log.Infof("NewSHM: SHMRaw.CloseVoteTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.CloseVoteTime), unsafe.Sizeof(Shm.Raw.CloseVoteTime))

    log.Infof("NewSHM: SHMRaw.Notes: (%v/%v)", unsafe.Offsetof(Shm.Raw.Notes), unsafe.Sizeof(Shm.Raw.Notes))

    log.Infof("NewSHM: SHMRaw.Gap18: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap18), unsafe.Sizeof(Shm.Raw.Gap18))

    log.Infof("NewSHM: SHMRaw.TodayIs: (%v/%v)", unsafe.Offsetof(Shm.Raw.TodayIs), unsafe.Sizeof(Shm.Raw.TodayIs))

    log.Infof("NewSHM: SHMRaw.NeverUsedNNotes_: (%v/%v)", unsafe.Offsetof(Shm.Raw.NeverUsedNNotes_), unsafe.Sizeof(Shm.Raw.NeverUsedNNotes_))

    log.Infof("NewSHM: SHMRaw.Gap19: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap19), unsafe.Sizeof(Shm.Raw.Gap19))

    log.Infof("NewSHM: SHMRaw.NeverUsedNextRefresh_: (%v/%v)", unsafe.Offsetof(Shm.Raw.NeverUsedNextRefresh_), unsafe.Sizeof(Shm.Raw.NeverUsedNextRefresh_))

    log.Infof("NewSHM: SHMRaw.Gap20: (%v/%v)", unsafe.Offsetof(Shm.Raw.Gap20), unsafe.Sizeof(Shm.Raw.Gap20))

    log.Infof("NewSHM: SHMRaw.LoginMsg: (%v/%v)", unsafe.Offsetof(Shm.Raw.LoginMsg), unsafe.Sizeof(Shm.Raw.LoginMsg))

    log.Infof("NewSHM: SHMRaw.LastFilm: (%v/%v)", unsafe.Offsetof(Shm.Raw.LastFilm), unsafe.Sizeof(Shm.Raw.LastFilm))

    log.Infof("NewSHM: SHMRaw.LastUsong: (%v/%v)", unsafe.Offsetof(Shm.Raw.LastUsong), unsafe.Sizeof(Shm.Raw.LastUsong))

    log.Infof("NewSHM: SHMRaw.PUptime: (%v/%v)", unsafe.Offsetof(Shm.Raw.PUptime), unsafe.Sizeof(Shm.Raw.PUptime))

    log.Infof("NewSHM: SHMRaw.PTouchTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.PTouchTime), unsafe.Sizeof(Shm.Raw.PTouchTime))

    log.Infof("NewSHM: SHMRaw.PBusyState: (%v/%v)", unsafe.Offsetof(Shm.Raw.PBusyState), unsafe.Sizeof(Shm.Raw.PBusyState))

    log.Infof("NewSHM: SHMRaw.GV2: (%v/%v)", unsafe.Offsetof(Shm.Raw.GV2), unsafe.Sizeof(Shm.Raw.GV2))

    log.Infof("NewSHM: SHMRaw.Statistic: (%v/%v)", unsafe.Offsetof(Shm.Raw.Statistic), unsafe.Sizeof(Shm.Raw.Statistic))

    log.Infof("NewSHM: SHMRaw.DeprecatedHomeIp_: (%v/%v)", unsafe.Offsetof(Shm.Raw.DeprecatedHomeIp_), unsafe.Sizeof(Shm.Raw.DeprecatedHomeIp_))

    log.Infof("NewSHM: SHMRaw.DeprecatedHomeMask_: (%v/%v)", unsafe.Offsetof(Shm.Raw.DeprecatedHomeMask_), unsafe.Sizeof(Shm.Raw.DeprecatedHomeMask_))

    log.Infof("NewSHM: SHMRaw.DeprecatedHomeDesc_: (%v/%v)", unsafe.Offsetof(Shm.Raw.DeprecatedHomeDesc_), unsafe.Sizeof(Shm.Raw.DeprecatedHomeDesc_))

    log.Infof("NewSHM: SHMRaw.DeprecatedHomeNum_: (%v/%v)", unsafe.Offsetof(Shm.Raw.DeprecatedHomeNum_), unsafe.Sizeof(Shm.Raw.DeprecatedHomeNum_))

    log.Infof("NewSHM: SHMRaw.MaxUser: (%v/%v)", unsafe.Offsetof(Shm.Raw.MaxUser), unsafe.Sizeof(Shm.Raw.MaxUser))

    log.Infof("NewSHM: SHMRaw.MaxTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.MaxTime), unsafe.Sizeof(Shm.Raw.MaxTime))

    log.Infof("NewSHM: SHMRaw.FUptime: (%v/%v)", unsafe.Offsetof(Shm.Raw.FUptime), unsafe.Sizeof(Shm.Raw.FUptime))

    log.Infof("NewSHM: SHMRaw.FTouchTime: (%v/%v)", unsafe.Offsetof(Shm.Raw.FTouchTime), unsafe.Sizeof(Shm.Raw.FTouchTime))

    log.Infof("NewSHM: SHMRaw.FBusyState: (%v/%v)", unsafe.Offsetof(Shm.Raw.FBusyState), unsafe.Sizeof(Shm.Raw.FBusyState))

    size := SHMSIZE

    if isCreate {
        shmid, shmaddr, isNew, err = shm.CreateShm(key, size, isUseHugeTlb)
        if err != nil {
            return err
        }
    } else {
        shmid, shmaddr, err = shm.OpenShm(key, size, isUseHugeTlb)
        if err != nil {
            return err
        }
    }

    Shm = &SHM{
        Shmid:   shmid,
        IsNew:   isNew,
        Shmaddr: shmaddr,
    }

    if isNew {
        in_version := SHM_VERSION
        in_size := int32(SHM_RAW_SZ)
        in_number := int32(0)
        in_loaded := int32(0)
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Raw.Version),
            unsafe.Sizeof(Shm.Raw.Version),
            unsafe.Pointer(&in_version),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Raw.Size),
            unsafe.Sizeof(Shm.Raw.Size),
            unsafe.Pointer(&in_size),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Raw.Number),
            unsafe.Sizeof(Shm.Raw.Number),
            unsafe.Pointer(&in_number),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Raw.Loaded),
            unsafe.Sizeof(Shm.Raw.Loaded),
            unsafe.Pointer(&in_loaded),
        )
    }

    // version and size should be fixed.
    Shm.ReadAt(
        unsafe.Offsetof(Shm.Raw.Version),
        unsafe.Sizeof(Shm.Raw.Version),
        unsafe.Pointer(&Shm.Raw.Version),
    )
    Shm.ReadAt(
        unsafe.Offsetof(Shm.Raw.Size),
        unsafe.Sizeof(Shm.Raw.Size),
        unsafe.Pointer(&Shm.Raw.Size),
    )

    // verify version
    if Shm.Raw.Version != SHM_VERSION {
        log.Errorf("NewSHM: version not match: key: %v Shm.Raw.Version: %v SHM_VERSION: %v isCreate: %v isNew: %v", key, Shm.Raw.Version, SHM_VERSION, isCreate, isNew)
        debug.PrintStack()
        CloseSHM()
        return ErrShmVersion
    }
    if Shm.Raw.Size != int32(SHM_RAW_SZ) {
        log.Warnf("NewSHM: size not match (version matched): key: %v Shm.Raw.Size: %v SHM_RAW_SZ: %v size: %v isCreate: %v isNew: %v", key, Shm.Raw.Size, SHM_RAW_SZ, size, isCreate, isNew)

        CloseSHM()
        return ErrShmSize
    }

    if isCreate && !isNew {
        log.Warnf("NewSHM: is expected to create, but not: key: %v", key)
    }

    log.Infof("NewSHM: shm created: key: %v size: %v isNew: %v", key, Shm.Raw.Size, isNew)

    return nil
}

//Close
//
//XXX [WARNING] know what you are doing before using Close!.
//This is to be able to close the shared mem for the completeness of the mem-usage.
//However, in production, we create shm without the need of closing the shm.
//
//We simply use ipcrm to delete the shm if necessary.
//
//Currently used only in test.
func CloseSHM() error {
    if Shm == nil {
        // Already Closed
        log.Errorf("CloseSHM: already closed")
        return ErrShmNotInit
    }

    err := Shm.Close()
    if err != nil {
        log.Errorf("CloseSHM: unable to close: e: %v", err)
        return err
    }

    Shm = nil

    log.Infof("CloseSHM: done")

    return nil
}

//Close
//
//XXX [WARNING] know what you are doing before using Close!.
//This is to be able to close the shared mem for the completeness of the mem-usage.
//However, in production, we create shm without the need of closing the shm.
//
//We simply use ipcrm to delete the shm if necessary.
//
//Currently used only in test.
func (s *SHM) Close() error {
    if !IsTest {
        return ErrInvalidOp
    }
    return shm.CloseShm(s.Shmid)
}

//ReadAt
//
//Require precalculated offset and size and outptr to efficiently get the data.
//See tests for exact usage.
//[!!!] If we are reading from the array, make sure that have unit-size * n in the size.
//
//Params
//  offsetOfSHMRawComponent: offset from SHMRaw
//  size: size of the variable, usually can be referred from SHMRaw
//        [!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//  outptr: the ptr of the object to read.
func (s *SHM) ReadAt(offsetOfSHMRawComponent uintptr, size uintptr, outptr unsafe.Pointer) {
    shm.ReadAt(s.Shmaddr, int(offsetOfSHMRawComponent), types.Size_t(size), outptr)
}

//WriteAt
//
//Require recalculated offset and size and outptr to efficiently get the data.
//See tests for exact usage.
//[!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//
//Params
//  offsetOfSHMRawComponent: offset from SHMRaw
//  size: size of the variable
//        [!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//  inptr: the ptr of the object to write.
func (s *SHM) WriteAt(offsetOfSHMRawComponent uintptr, size uintptr, inptr unsafe.Pointer) {
    shm.WriteAt(s.Shmaddr, int(offsetOfSHMRawComponent), types.Size_t(size), inptr)
}

func (s *SHM) SetOrUint32(offsetOfSHMRawComponent uintptr, inptr unsafe.Pointer) {
    shm.SetOrUint32(s.Shmaddr, int(offsetOfSHMRawComponent), inptr)
}

func (s *SHM) IncUint32(offsetOfSHMRawComponent uintptr) {
    shm.IncUint32(s.Shmaddr, int(offsetOfSHMRawComponent))
}

func (s *SHM) Cmp(offsetOfSHMRawComponent uintptr, size uintptr, cmpptr unsafe.Pointer) int {
    return shm.Cmp(s.Shmaddr, int(offsetOfSHMRawComponent), types.Size_t(size), cmpptr)
}
