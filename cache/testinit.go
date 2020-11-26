package cache

import (
	"sync"

	"github.com/PichuChen/go-bbs/ptttype"
)

const (
	TestShmKey = 2000000
)

var (
	IsTest = false

	origBBSHome = ""

	testMutex sync.Mutex
)

func setupTest() {
	testMutex.Lock()
	IsTest = true
	origBBSHome = ptttype.SetBBSHOME("./testcase")
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHome)
	IsTest = false
	testMutex.Unlock()
}
