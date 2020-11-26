package cmbbs

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	log "github.com/sirupsen/logrus"
)

var (
	origBBSHOME = ""
)

func setupTest() {
	cache.TestMutex.Lock()
	origBBSHOME = ptttype.SetBBSHOME("./testcase")

	err := cache.NewSHM(cache.TestShmKey, ptttype.USE_HUGETLB, true)
	if err != nil {
		log.Errorf("setupTest: unable to NewSHM: e: %v", err)
		return
	}
	_ = cache.LoadUHash()
}

func teardownTest() {
	_ = cache.CloseSHM()
	ptttype.SetBBSHOME(origBBSHOME)
	cache.TestMutex.Unlock()
}
