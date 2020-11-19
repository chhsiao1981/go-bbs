package main

import (
	"flag"

	"github.com/PichuChen/go-bbs/api"
	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
)

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	router.POST("/login", NewApi(api.Login, &api.LoginParams{}).Json)
	router.POST("/ping", NewApi(api.Ping, nil).LoginRequiredJson)

	return router, nil
}

func initMain() error {
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)
	log.SetLevel(log.InfoLevel)

	filename := ""
	flag.StringVar(&filename, "ini", "config.ini", "ini filename")
	flag.Parse()

	_ = InitConfig(filename)

	err := cache.NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, ptttype.IS_NEW_SHM)
	if err != nil {
		log.Errorf("unable to init SHM: e: %v", err)
		return err
	}

	if ptttype.IS_NEW_SHM {
		err = cache.LoadUHash()
		if err != nil {
			log.Errorf("unable to load UHash: e: %v", err)
			return err
		}
	}
	err = cache.AttachCheckSHM()
	if err != nil {
		log.Errorf("unable to attach-check-shm: e: %v", err)
		return err
	}

	return nil
}

func main() {
	err := initMain()
	if err != nil {
		log.Errorf("unable to initMain: e: %v", err)
		return
	}
	router, err := initGin()
	if err != nil {
		return
	}
	router.Run(HTTP_HOST)
}
