package ptt

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

func Test_registerCountEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		user  *ptttype.UserecRaw
		email *[ptttype.EMAILSZ]byte
	}
	tests := []struct {
		name          string
		args          args
		expectedCount int
		wantErr       bool
	}{
		// TODO: Add test cases.
		{},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := registerCountEmail(tt.args.user, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerCountEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.expectedCount {
				t.Errorf("registerCountEmail() = %v, expected %v", gotCount, tt.expectedCount)
			}
		})
	}
}

func Test_getSystemUaVersion(t *testing.T) {
	tests := []struct {
		name     string
		expected uint8
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSystemUaVersion(); got != tt.expected {
				t.Errorf("getSystemUaVersion() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestSetupNewUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		user *ptttype.UserecRaw
	}
	tests := []struct {
		name     string
		args     args
		expected *ptttype.UserecRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{testSetupNewUser1},
			expected: testSetupNewUser1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetupNewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SetupNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_, got, err := initCurrentUser(&tt.args.user.UserID)
			if err != nil {
				t.Errorf("SetupNewUser (initCurrentUser): err: %v", err)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SetupNewUser (initCurrentUser): got: %v expected: %v", got, tt.expected)
			}
		})
	}
}

func Test_isToCleanUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	file, err := os.OpenFile(ptttype.FN_FRESH, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Warnf("unable to open-file: e: %v", err)
	}
	file.Write([]byte("temp"))
	file.Close()

	newTime1 := time.Now().Add(-3700 * types.TS_TO_NANO_TS)
	newTime2 := time.Now().Add(-2700 * types.TS_TO_NANO_TS)

	tests := []struct {
		name     string
		newTime  time.Time
		isDelete bool
		expected bool
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "old-file",
			newTime:  newTime1,
			expected: true,
		},
		{
			name:     "new-file",
			newTime:  newTime2,
			expected: false,
		},
		{
			name:     "no file",
			isDelete: true,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isDelete {
				os.Remove(ptttype.FN_FRESH)
			} else {
				err = os.Chtimes(ptttype.FN_FRESH, tt.newTime, tt.newTime)
				if err != nil {
					t.Errorf("unable to Chtimes e: %v", err)
				}
			}

			got, err := isToCleanUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("isToCleanUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("isToCleanUser() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func Test_touchFresh(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := touchFresh(); (err != nil) != tt.wantErr {
				t.Errorf("touchFresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkAndExpireAccount(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid         int32
		user        *ptttype.UserecRaw
		expireRange int
	}
	tests := []struct {
		name     string
		args     args
		expected int
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkAndExpireAccount(tt.args.uid, tt.args.user, tt.args.expireRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkAndExpireAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("checkAndExpireAccount() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func Test_computeUserExpireValue(t *testing.T) {
	setupTest()
	defer teardownTest()

	user1 := &ptttype.UserecRaw{}
	user1.UserLevel |= ptttype.PERM_XEMPT

	user2 := &ptttype.UserecRaw{}
	copy(user2.UserID[:], ptttype.USER_ID_REGNEW[:])
	user2.LastLogin = types.NowTS() - 10*60

	user3 := &ptttype.UserecRaw{}
	copy(user3.UserID[:], ptttype.USER_ID_REGNEW[:])
	user3.LastLogin = types.NowTS() - 400*60

	user4 := &ptttype.UserecRaw{}
	user4.LastLogin = types.NowTS() - 119*24*60*60

	type args struct {
		user *ptttype.UserecRaw
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		// TODO: Add test cases.
		{
			name:     "PERM_XEMPT",
			args:     args{user1},
			expected: true,
		},
		{
			name:     "new (valid)",
			args:     args{user2},
			expected: true,
		},
		{
			name:     "new (invalid)",
			args:     args{user3},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeUserExpireValue(tt.args.user)
			if got < 0 && tt.expected {
				t.Errorf("computeUserExpireValue() = %v, expected %v", got, tt.expected)
			} else if got >= 0 && !tt.expected {
				t.Errorf("computeUserExpireValue() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestNewRegister(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userID          *[ptttype.IDLEN + 1]byte
		passwd          []byte
		fromHost        *[ptttype.IPV4LEN + 1]byte
		email           *[ptttype.EMAILSZ]byte
		isEmailVerified bool
		isAdbannerUSong bool
		nickname        *[ptttype.NICKNAMESZ]byte
		realname        *[ptttype.REALNAMESZ]byte
		career          *[ptttype.CAREERSZ]byte
		address         *[ptttype.ADDRESSSZ]byte
		over18          bool
	}
	tests := []struct {
		name     string
		args     args
		expected *ptttype.UserecRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userID:   &testNewRegister1.UserID,
				passwd:   testNewRegister1Passwd,
				fromHost: &testNewRegister1.LastHost,
				email:    &testNewRegister1.Email,
				nickname: &testNewRegister1.Nickname,
				realname: &testNewRegister1.RealName,
				career:   &testNewRegister1.Career,
				address:  &testNewRegister1.Address,
				over18:   testNewRegister1.Over18,
			},
			expected: testNewRegister1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRegister(tt.args.userID, tt.args.passwd, tt.args.fromHost, tt.args.email, tt.args.isEmailVerified, tt.args.isAdbannerUSong, tt.args.nickname, tt.args.realname, tt.args.career, tt.args.address, tt.args.over18)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			isGood, err := cmbbs.CheckPasswd(got.PasswdHash[:], testNewRegister1Passwd)
			if err != nil || !isGood {
				t.Errorf("NewRegister() unable to checkpasswd: passwd: %v", string(testNewRegister1Passwd))
			}
			copy(testNewRegister1.PasswdHash[:], got.PasswdHash[:])
			testNewRegister1.LastLogin = got.LastLogin
			testNewRegister1.FirstLogin = got.FirstLogin
			testNewRegister1.LastSeen = got.LastSeen

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewRegister() = %v, want %v", got, tt.expected)
			}
		})
	}
}
