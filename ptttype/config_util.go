package ptttype

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	log.Infof("InitConfig: start")

	initConfig()

	postInitConfig()
}

func setStringConfig(idx string, orig string) string {
	idx = "ptttype." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetString(idx)
}

func setBoolConfig(idx string, orig bool) bool {
	idx = "ptttype." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetBool(idx)
}

func setColorConfig(idx string, orig string) string {
	idx = "ptttype." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}
	return ANSIColor(viper.GetString(idx))
}

func setIntConfig(idx string, orig int) int {
	idx = "ptttype." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}
	return viper.GetInt(idx)
}

func setDoubleConfig(idx string, orig float64) float64 {
	idx = "ptttype." + strings.ToLower(idx)
	if !viper.IsSet(idx) {
		return orig
	}

	return viper.GetFloat64(idx)
}

//SetBBSHOME
//
//This is to safely set BBSHOME
//Public to be used in the tests of other modules.
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

//setBBSMName
//
//This is to safely set BBSMNAME
//
//Params
//	bbsmname: new bbsmname
//
//Return
//	string: original bbsmname
func setBBSMName(bbsmname string) string {
	origBBSMName := BBSMNAME
	log.Debugf("SetBBSMNAME: %v", bbsmname)

	BBSMNAME = bbsmname

	// regex-replace

	BN_SECURITY = regexReplace(BN_SECURITY, "BBSMNAME", BBSMNAME)
	BN_NOTE = regexReplace(BN_NOTE, "BBSMNAME", BBSMNAME)
	BN_RECORD = regexReplace(BN_RECORD, "BBSMNAME", BBSMNAME)
	BN_SYSOP = regexReplace(BN_SYSOP, "BBSMNAME", BBSMNAME)
	BN_TEST = regexReplace(BN_SECURITY, "BBSMNAME", BBSMNAME)
	BN_BUGREPORT = regexReplace(BN_BUGREPORT, "BBSMNAME", BBSMNAME)
	BN_LAW = regexReplace(BN_LAW, "BBSMNAME", BBSMNAME)
	BN_NEWBIE = regexReplace(BN_NEWBIE, "BBSMNAME", BBSMNAME)
	BN_ASKBOARD = regexReplace(BN_ASKBOARD, "BBSMNAME", BBSMNAME)
	BN_FOREIGN = regexReplace(BN_FOREIGN, "BBSMNAME", BBSMNAME)

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

	BN_BUGREPORT = BBSMNAME + "Bug"
	BN_LAW = BBSMNAME + "Law"
	BN_NEWBIE = BBSMNAME + "NewHand"
	BN_FOREIGN = BBSMNAME + "Foreign"

	return origBBSMName
}

func regexReplace(str string, substr string, repl string) string {
	theRe := regexp.MustCompile("\\s*" + substr + "\\s*")
	if theRe == nil {
		return str
	}

	return theRe.ReplaceAllString(str, repl)
}

func setCAPTCHAInsertServerAddr(captchaInsertServerAddr string) string {
	origCAPTCHAInsertServerAddr := CAPTCHA_INSERT_SERVER_ADDR

	CAPTCHA_INSERT_SERVER_ADDR = captchaInsertServerAddr

	if IS_CAPTCHA_INSERT_HOST_INFERRED {
		CAPTCHA_INSERT_HOST = CAPTCHA_INSERT_SERVER_ADDR
	}

	return origCAPTCHAInsertServerAddr
}

//setMyHostname
//
//Params
//	myHostName: new my hostname
//
//Return
//	string: orig my hostname
func setMyHostname(myHostname string) string {
	origMyHostname := MYHOSTNAME

	MYHOSTNAME = myHostname

	if IS_AID_HOSTNAME_INFERRED {
		AID_HOSTNAME = MYHOSTNAME
	}

	return origMyHostname

}

//setRecycleBinName
//
//Params
//	recycleBinName: new recycle bin name
//
//Return
//	string: orig recycle bin name
func setRecycleBinName(recycleBinName string) string {
	origRecycleBinName := recycleBinName

	RECYCLE_BIN_NAME = recycleBinName
	RECYCLE_BIN_OWNER = "[" + RECYCLE_BIN_NAME + "]"

	return origRecycleBinName
}

func postInitConfig() {
	log.Infof("postInitConfig: start: BBSHOME: %v", BBSHOME)
	_ = SetBBSHOME(BBSHOME)
	_ = setBBSMName(BBSMNAME)
	_ = setCAPTCHAInsertServerAddr(CAPTCHA_INSERT_SERVER_ADDR)
	_ = setMyHostname(MYHOSTNAME)
	_ = setRecycleBinName(RECYCLE_BIN_NAME)
}
