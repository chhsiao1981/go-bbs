package bbs

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestNewUserecFromRaw(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userecraw *ptttype.UserecRaw
	}
	tests := []struct {
		name     string
		args     args
		expected *Userec
	}{
		// TODO: Add test cases.
		{
			args:     args{testUserecRaw},
			expected: testUserec1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserecFromRaw(tt.args.userecraw); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewUserecFromRaw() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
