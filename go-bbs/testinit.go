package main

import (
	"os"
	"sync"

	"github.com/PichuChen/go-bbs/cache"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	"github.com/gin-gonic/gin"

	jww "github.com/spf13/jwalterweatherman"
)

var (
	testMutex sync.Mutex
)

func setupTest() {
	testMutex.Lock()
	jww.SetLogOutput(os.Stderr)
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)

	initConfig("./testcase/test.ini")

	gin.SetMode(gin.TestMode)

	// shm
	cache.IsTest = true
	_ = cache.NewSHM(types.Key_t(cache.TestShmKey), ptttype.USE_HUGETLB, true)

	cache.LoadUHash()
	cache.AttachSHM()
}

func teardownTest() {
	cache.CloseSHM()
	cache.IsTest = false
	testMutex.Unlock()
}
