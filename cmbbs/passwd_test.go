package cmbbs

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestPasswdLoadUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID1[:], []byte("SYSOP"))

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
			args:      args{&userID1},
			expected:  1,
			expected1: testUserecRaw1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PasswdLoadUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswdLoadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("PasswdLoadUser() got = %v, expected %v", got, tt.expected)
			}
			if !reflect.DeepEqual(got1, tt.expected1) {
				t.Errorf("PasswdLoadUser() got1 = %v, expected1 %v", got1, tt.expected1)
			}
		})
	}
}

func TestPasswdQuery(t *testing.T) {
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
			got, err := PasswdQuery(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswdQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("PasswdQuery() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCheckPasswd(t *testing.T) {
	input1 := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', 0}
	input2 := []byte{'0', '1', '2', '4', '4', '5', '6', '7', '8', '9', '0', '1', 0}
	expected1 := [ptttype.PASSLEN]byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65}

	type args struct {
		expected []byte
		input    []byte
	}
	tests := []struct {
		name     string
		args     args
		expected bool
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{expected1[:], input1},
			expected: true,
		},
		{
			name:     "incorrect input",
			args:     args{expected1[:], input2},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckPasswd(tt.args.expected, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("CheckPasswd() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestGenPasswd(t *testing.T) {
	type args struct {
		passwd []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{[]byte("1234")},
		},
		{
			args: args{[]byte("abcdef")},
		},
		{
			args: args{[]byte("834792134")},
		},
		{
			args: args{[]byte("rweqrrwe")},
		},
		{
			args: args{[]byte("!@#$5ks")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPasswdHash, err := GenPasswd(tt.args.passwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenPasswd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			isGood, err := CheckPasswd(gotPasswdHash[:], tt.args.passwd)
			if err != nil || !isGood {
				t.Errorf("GenPasswd: unable to pass CheckPasswd: passwd: %v gotPasswdHash: %v", tt.args.passwd, gotPasswdHash)
			}
		})
	}
}
