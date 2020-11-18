package ptt

import (
	"os"
	"sync"

	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

var (
	origBBSHOME = ""

	testMutex sync.Mutex
)

func setupTest() {
	testMutex.Lock()
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
	testMutex.Unlock()
}
