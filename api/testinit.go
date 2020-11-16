package api

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

var (
	testOrigBBSHOME = ""
)

func setupTest() {
	testOrigBBSHOME = ptttype.SetBBSHOME("./testcase")

	// shm
	_ = cache.NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)
	cache.LoadUHash()
	cache.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(testOrigBBSHOME)
	cache.CloseSHM()
}
