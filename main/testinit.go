package main

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
	"github.com/gin-gonic/gin"
)

var (
	testOrigBBSHOME = ""
)

func setupTest() {
	testOrigBBSHOME = ptttype.SetBBSHOME("./testcase")

	gin.SetMode(gin.TestMode)

	// shm
	shm.LoadUHash()
	shm.AttachSHM()
}

func teardownTest() {
	ptttype.SetBBSHOME(testOrigBBSHOME)
}
