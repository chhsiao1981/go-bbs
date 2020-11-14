package ptttype

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(filename string) error {

	filenameList := strings.Split(filename, ".")
	if len(filenameList) == 1 {
		return ErrInvalidIni
	}

	filenamePrefix := strings.Join(filenameList[:len(filenameList)-1], ".")
	filenamePostfix := filenameList[len(filenameList)-1]
	viper.SetConfigName(filenamePrefix)
	viper.SetConfigType(filenamePostfix)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	log.Infof("viper keys: %v", viper.AllKeys())

	initConfig()
	return nil
}

func setStringConfig(idx string, orig string) string {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}

func setBoolConfig(idx string, orig bool) bool {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetBool(idx)
}

func setColorConfig(idx string, orig string) string {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}
	return ANSIColor(viper.GetString(idx))
}

func setIntConfig(idx string, orig int) int {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}
	return viper.GetInt(idx)
}

func setDoubleConfig(idx string, orig float64) float64 {
	idx = "go-bbs." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetFloat64(idx)
}

//SetBBSHOME
//
//This is to safely set BBSHOME
//
//Params
//	bbshome: new bbshome
//
//Return
//	string: original bbshome
func SetBBSHOME(bbshome string) string {
	origBBSHome := BBSHOME
	log.Debugf("SetBBSHOME: %v", bbshome)

	// config.go
	BBSHOME = bbshome
	BBSPROG = BBSHOME + BBSPROGPOSTFIX

	//common.go
	FN_CONF_BANIP = BBSHOME + FN_CONF_BANIP_POSTFIX // 禁止連線的 IP 列表
	FN_PASSWD = BBSHOME + FN_PASSWD_POSTFIX         /* User records */

	return origBBSHome
}

//SetBBSMNAME
//
//This is to safely set BBSMNAME
//
//Params
//	bbsmname: new bbsmname
//
//Return
//	string: original bbsmname
func SetBBSMNAME(bbsmname string) string {
	origBBSMName := BBSMNAME
	log.Debugf("SetBBSMNAME: %v", bbsmname)

	BBSMNAME = bbsmname

	// config.go
	if IS_BN_FIVECHESS_LOG_INFERRED {
		BN_FIVECHESS_LOG = BBSMNAME + "Five"
	}
	if IS_BN_CCHESS_LOG_INFERRED {
		BN_CCHESS_LOG = BBSMNAME + "CChess"
	}
	if IS_MONEYNAME_INFFERRED {
		MONEYNAME = BBSMNAME + "幣"
	}

	//common.go

	return origBBSMName
}
