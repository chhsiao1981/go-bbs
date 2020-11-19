package main

import (
	"os"

	"github.com/PichuChen/go-bbs/cache"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	"github.com/gin-gonic/gin"

	jww "github.com/spf13/jwalterweatherman"
)

var ()

func setupTest() {
	jww.SetLogOutput(os.Stderr)
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)

	_ = InitConfig("./testcase/test.ini")

	gin.SetMode(gin.TestMode)

	// shm
	cache.IsTest = true
	_ = cache.NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)

	_ = cache.LoadUHash()
	_ = cache.AttachSHM()
}

func teardownTest() {
	_ = cache.CloseSHM()
	cache.IsTest = false
}
