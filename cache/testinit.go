package cache

import (
	"sync"
	"time"

	"github.com/PichuChen/go-bbs/ptttype"
)

const (
	TestShmKey = 1000000
)

var (
	IsTest = false

	origBBSHome = ""
	origShmKey  = 0

	TestMutex sync.Mutex
)

func setupTest() {
	TestMutex.Lock()
	IsTest = true
	origShmKey = ptttype.SHM_KEY
	ptttype.SHM_KEY = TestShmKey
	origBBSHome = ptttype.SetBBSHOME("./testcase")
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHome)
	IsTest = false
	ptttype.SHM_KEY = origShmKey
	TestMutex.Unlock()
	time.Sleep(5000000)
}
