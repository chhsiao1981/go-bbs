package ptt

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/PichuChen/go-bbs/cache"
	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

//NewRegister
//
//XXX Assuming valid input. need to verify email at api.
func NewRegister(
	userID *[ptttype.IDLEN + 1]byte,
	passwd []byte,
	fromHost *[ptttype.IPV4LEN + 1]byte,
	email *[ptttype.EMAILSZ]byte,
	isEmailVerified bool,
	isAdbannerUSong bool,

	nickname *[ptttype.NICKNAMESZ]byte,
	realname *[ptttype.REALNAMESZ]byte,
	career *[ptttype.CAREERSZ]byte,
	address *[ptttype.ADDRESSSZ]byte,
	over18 bool,
) (*ptttype.UserecRaw, error) {

	// line: 758
	newUser := &ptttype.UserecRaw{}
	newUser.Version = ptttype.PASSWD_VERSION
	newUser.UserLevel = ptttype.PERM_DEFAULT
	newUser.UFlag = ptttype.UF_BRDSORT | ptttype.UF_ADBANNER | ptttype.UF_CURSOR_ASCII
	newUser.FirstLogin = types.NowTS()
	newUser.LastLogin = newUser.FirstLogin
	newUser.Pager = ptttype.PAGER_ON
	newUser.NumLoginDays = 1
	newUser.UaVersion = getSystemUaVersion()
	copy(newUser.LastHost[:], fromHost[:])

	if ptttype.DBCSAWARE {
		newUser.UFlag |= ptttype.UF_DBCS_AWARE | ptttype.UF_DBCS_DROP_REPEAT
	}

	// XXX REQUIRE_VERIFY_EMAIL_AT_REGISTER
	if email != nil {
		copy(newUser.Email[:], email[:])
	}

	// XXX UF_ADBANNER_USONG
	if isAdbannerUSong {
		newUser.UFlag |= ptttype.UF_ADBANNER_USONG
	}

	// line: 857
	passwdHash, err := cmbbs.GenPasswd(passwd)
	if err != nil {
		return nil, err
	}
	copy(newUser.PasswdHash[:], passwdHash[:])

	copy(newUser.Nickname[:], nickname[:])
	copy(newUser.RealName[:], realname[:])
	copy(newUser.Career[:], career[:])
	copy(newUser.Address[:], address[:])
	newUser.Over18 = over18

	// XXX SECURE_LOGIN
	// XXX no need to secure-login for now.
	// newUser.UFlag |= ptttype.UF_SECURE_LOGIN

	copy(newUser.UserID[:], userID[:])

	err = SetupNewUser(newUser)
	if err != nil {
		//log.Errorf("register.NewRegister: unable to SetupNewUser: userID: %v e: %v", userID, err)
		return nil, err
	}

	// read and ensure that every works as expected.
	uid, _, err := initCurrentUser(userID)
	if err != nil {
		log.Errorf("register.NewRegister: unable to initCurrentUser: userID: %v e: %v", userID, err)
		return nil, err
	}

	err = ensureErasingOldUser(uid, userID)
	if err != nil {
		return nil, err
	}

	// XXX no define of USE_REMOVEBM_ON_NEWREG (line: 985)

	// if email verified.
	if isEmailVerified {
		emailErr := registerCheckAndUpdateEmaildb(newUser, &newUser.Email)
		if emailErr == nil {
			justify := [ptttype.REGLEN + 1]byte{}
			copy(justify[:ptttype.REGLEN], []byte(fmt.Sprintf("<E-Mail>: %v", types.NowTS().Cdate())))
			err = pwcuRegCompleteJustify(uid, userID, &justify)
			if err != nil {
				return nil, err
			}
		}
	}

	user, err := passwdSyncQuery(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ensureErasingOldUser(uid int32, userID *[ptttype.IDLEN + 1]byte) (err error) {
	filename := cmbbs.SetHomePath(userID)
	tmpFilename := filename + fmt.Sprintf(".%v", types.NowTS())
	if !types.IsDir(filename) {
		return nil
	}

	err = os.Rename(filename, tmpFilename)
	if err == nil {
		return nil
	}

	pwcuBitDisableLevel(uid, userID, ptttype.PERM_BASIC)

	return nil
}

func registerCheckAndUpdateEmaildb(user *ptttype.UserecRaw, email *[ptttype.EMAILSZ]byte) (err error) {

	_, err = registerCountEmail(user, email)
	if err != nil {
		return err
	}

	if ptttype.USE_EMAILDB {
		err = emailDBUpdateEmail(&user.UserID, email)
		if err != nil {
			return err
		}
	}

	if ptttype.USE_VERIFYDB {
		lcemail := bytes.ToLower(email[:])
		err = verifyDBSet(&user.UserID, int64(user.FirstLogin), ptttype.VMETHOD_EMAIL, lcemail, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func registerCountEmail(user *ptttype.UserecRaw, email *[ptttype.EMAILSZ]byte) (count int, err error) {

	if ptttype.USE_EMAILDB {
		r, err := emailDBCheckEmail(&user.UserID, email)
		if err != nil {
			return r, err
		}
		count = r
	}

	if ptttype.USE_VERIFYDB {
		lcemail := bytes.ToLower(email[:])
		_, countOther, err := verifyDBCountByVerify(ptttype.VMETHOD_EMAIL, lcemail)
		if err != nil {
			return -1, err
		}

		if count < countOther {
			count = countOther
		}
	}

	return count, nil
}

func getSystemUaVersion() uint8 {
	file, err := os.Open(ptttype.HAVE_USERAGREEMENT_VERSION)
	if err != nil {
		return 0
	}
	defer file.Close()

	var version uint
	_, err = fmt.Fscanf(file, "%u", &version)
	if err != nil {
		return 0
	}

	if version > 255 {
		version = 255
	}

	return uint8(version)
}

func SetupNewUser(user *ptttype.UserecRaw) error {
	// XXX race-condition, need to setup RWLock (across-processes)
	//
	uid, err := cache.DoSearchUserRaw(&user.UserID, nil)
	if err != nil {
		return err
	}

	if uid != 0 {
		return ptttype.ErrUserIDAlreadyExists
	}

	/* Lazy method : 先找尋已經清除的過期帳號 */
	uid, err = cache.DoSearchUserRaw(&ptttype.EMPTY_USER_ID, nil)
	if err != nil {
		return err
	}

	if uid == 0 { //unable to find empty user.
		err := tryCleanUser()
		if err != nil {
			return err
		}
	}

	//init passwd-semaphores
	err = cmbbs.PasswdInit()
	if err != nil {
		return err
	}

	err = cmbbs.PasswdLock()
	if err != nil {
		return err
	}
	defer cmbbs.PasswdUnlock()

	uid, err = cache.DoSearchUserRaw(&ptttype.EMPTY_USER_ID, nil)
	if err != nil {
		return err
	}

	err = cache.SetUserID(uid, &user.UserID)
	if err != nil {
		return err
	}

	_, _ = cache.SetUMoney(uid, user.Money)

	err = passwdSyncUpdate(uid, user)
	if err != nil {
		return err
	}

	return nil
}

func tryCleanUser() error {

	isToCleanUser, err := isToCleanUser()
	if err != nil {
		return err
	}
	if !isToCleanUser {
		return nil
	}

	err = touchFresh()
	if err != nil {
		return err
	}

	/* 不曉得為什麼要從 2 開始... Ptt:因為SYSOP在1 */
	for uid := int32(2); uid <= ptttype.MAX_USERS; uid++ {
		// XXX ignoring err, do the log
		user, err := passwdSyncQuery(uid)
		if err != nil {
			log.Errorf("register.tryCleanUser: unable to PasswdSyncQuery: uid: %v e: %v", uid, err)
		}

		_, err = checkAndExpireAccount(uid, user, ptttype.CLEAN_USER_EXPIRE_RANGE_MIN)
		if err != nil {
			log.Errorf("register.tryCleanUser: unable to checkAndExpireAccount: uid: %v e: %v", uid, err)
		}

	}

	return nil
}

//isToCleanUser
func isToCleanUser() (bool, error) {
	theStat, err := os.Stat(ptttype.FN_FRESH)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return true, nil
		}
		return false, err
	}

	return theStat.ModTime().Before(time.Now().Add(-3600 * types.TS_TO_NANO_TS)), nil
}

//touchFresh
func touchFresh() error {
	file, err := os.OpenFile(ptttype.FN_FRESH, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(time.Now().String())

	return nil
}

func checkAndExpireAccount(uid int32, user *ptttype.UserecRaw, expireRange int) (int, error) {
	expireValue := computeUserExpireValue(user)
	if expireValue >= 0 { // not expired yet.
		return expireValue, nil
	}

	if -expireValue > expireRange {
		killUser(uid, &user.UserID)
	} else {
		expireValue = 0
	}

	return expireValue, nil
}

func computeUserExpireValue(user *ptttype.UserecRaw) int {
	if (user.UserID[0] == byte(0)) ||
		(user.UserLevel|ptttype.PERM_XEMPT) != 0 ||
		bytes.Equal(user.UserID[:], ptttype.USER_ID_GUEST[:]) { // no expire
		return 999999
	}

	valMinute := int(types.NowTS()-user.LastLogin) / 60 // min-time since last login

	if bytes.Equal(user.UserID[:], ptttype.USER_ID_REGNEW[:]) { // allow only 30 minutes when doing new-user.
		return 30 - valMinute
	}

	if user.UserLevel&(ptttype.PERM_LOGINOK|ptttype.PERM_VIOLATELAW) != 0 {
		return ptttype.KEEP_DAYS_REGGED*24*60 - valMinute
	}

	return ptttype.KEEP_DAYS_UNREGGED*24*60 - valMinute
}
