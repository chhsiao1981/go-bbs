package ptt

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
)

var (
	origBBSHOME = ""
)

func setupTest() {
	origBBSHOME = ptttype.SetBBSHOME("./testcase")
	shm.LoadUHash()
	shm.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHOME)
}
