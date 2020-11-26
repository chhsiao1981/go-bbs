package ptt

import (
	"os"

	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

var (
	origBBSHOME = ""
)

func setupTest() {
	cache.TestMutex.Lock()
	origBBSHOME = ptttype.SetBBSHOME("./testcase")
	_ = cache.NewSHM(types.Key_t(cache.TestShmKey), ptttype.USE_HUGETLB, true)

	_ = types.CopyFileToFile("./testcase/.PASSWDS1", "./testcase/.PASSWDS")

	_ = types.CopyDirToDir("./testcase/home1", "./testcase/home")

	_ = cache.LoadUHash()
	_ = cache.AttachSHM()
}

func teardownTest() {
	os.Remove("./testcase/.PASSWDS")
	os.RemoveAll("./testcase/home")
	ptttype.SetBBSHOME(origBBSHOME)
	_ = cache.CloseSHM()
	cache.TestMutex.Unlock()
}
