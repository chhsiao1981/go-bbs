package main

import (
	"github.com/PichuChen/go-bbs/api"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
	"github.com/gin-gonic/gin"
)

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	router.POST("/login", NewApi(api.Login, &api.LoginParams{}).Json)
	router.POST("/ping", NewApi(api.Ping, nil).LoginRequiredJson)

	return router, nil
}

func init() {
	ptttype.SetBBSHOME("./testcase")
	shm.LoadUHash()
	shm.AttachSHM()
}

func main() {

	router, err := initGin()
	if err != nil {
		return
	}
	router.Run(HTTP_HOST)
}
