package bbs

import (
	"testing"
)

func Test_setuserfile(t *testing.T) {
	type args struct {
		userID  string
		postfix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"testUserID", "testPostfix"},
			want: "/home/bbs/home/t/testUserID/testPostfix",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setuserfile(tt.args.userID, tt.args.postfix); got != tt.want {
				t.Errorf("setuserfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUserHomeDir(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"testUserID"},
			want: "/home/bbs/home/t/testUserID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserHomeDir(tt.args.userID); got != tt.want {
				t.Errorf("getUserHomeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
