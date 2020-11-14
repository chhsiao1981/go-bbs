package main

import (
	"os"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
	"github.com/gin-gonic/gin"

	jww "github.com/spf13/jwalterweatherman"
)

var ()

func setupTest() {
	jww.SetLogOutput(os.Stderr)
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)

	ptttype.InitConfig("./testcase/test.ini")

	gin.SetMode(gin.TestMode)

	// shm
	shm.LoadUHash()
	shm.AttachSHM()
}

func teardownTest() {
}
