package ptt

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

var (
	origBBSHOME = ""
)

func setupTest() {
	origBBSHOME = ptttype.SetBBSHOME("./testcase")
	_ = cache.NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)

	cache.LoadUHash()
	cache.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHOME)
	cache.CloseSHM()
}
