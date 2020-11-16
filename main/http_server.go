package main

import (
	"github.com/PichuChen/go-bbs/api"
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	router.POST("/login", NewApi(api.Login, &api.LoginParams{}).Json)
	router.POST("/ping", NewApi(api.Ping, nil).LoginRequiredJson)

	return router, nil
}

func initMain() error {
	ptttype.SetBBSHOME("./testcase")
	err := cache.NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)
	if err != nil {
		log.Errorf("unable to init SHM: e: %v", err)
		return err
	}
	cache.LoadUHash()
	cache.AttachSHM()

	return nil
}

func main() {

	router, err := initGin()
	if err != nil {
		return
	}
	router.Run(HTTP_HOST)
}
