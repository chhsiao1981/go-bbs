package api

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
)

var (
	testOrigBBSHOME = ""
)

func setupTest() {
	testOrigBBSHOME = ptttype.SetBBSHOME("./testcase")

	// shm
	shm.LoadUHash()
	shm.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(testOrigBBSHOME)
}
