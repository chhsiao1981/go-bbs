package cmbbs

import (
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
)

var (
	origBBSHOME = ""
)

func setupTest() {
	origBBSHOME = ptttype.SetBBSHOME("./testcase")
	cache.LoadUHash()
	cache.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHOME)
}
