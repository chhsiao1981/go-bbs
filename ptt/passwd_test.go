package ptt

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func Test_initCurrentUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	userid1 := [ptttype.IDLEN + 1]byte{}
	copy(userid1[:], []byte("SYSOP"))

	type args struct {
		userID *[ptttype.IDLEN + 1]byte
	}
	tests := []struct {
		name      string
		args      args
		expected  int32
		expected1 *ptttype.UserecRaw
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{&userid1},
			expected:  1,
			expected1: testUserecRaw1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := initCurrentUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("initCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("initCurrentUser() got = %v, expected %v", got, tt.expected)
			}
			if !reflect.DeepEqual(got1, tt.expected1) {
				t.Errorf("initCurrentUser() got1 = %v, expected1 %v", got1, tt.expected1)
			}
		})
	}
}

func Test_passwdSyncUpdate(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid  int32
		user *ptttype.UserecRaw
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{uid: int32(1), user: testUserecRaw1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := passwdSyncUpdate(tt.args.uid, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("passwdSyncUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_passwdSyncQuery(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid int32
	}
	tests := []struct {
		name     string
		args     args
		expected *ptttype.UserecRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{1},
			expected: testUserecRaw1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := passwdSyncQuery(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("passwdSyncQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("passwdSyncQuery() = %v, want %v", got, tt.expected)
			}
		})
	}
}
