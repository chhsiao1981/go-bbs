package main

import (
	"flag"

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

func initMain() {
	filename := ""
	flag.StringVar(&filename, "ini", "config.ini", "ini filename")
	flag.Parse()

	ptttype.InitConfig(filename)

	shm.LoadUHash()
	shm.AttachSHM()
}

func main() {
	initMain()
	router, err := initGin()
	if err != nil {
		return
	}
	router.Run(HTTP_HOST)
}
