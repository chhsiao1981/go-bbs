package ptttype

import (
	"testing"
)

func Test_regexReplace(t *testing.T) {
	type args struct {
		str    string
		substr string
		repl   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{str: "BBSMNAME test1", substr: "BBSMNAME", repl: "bbs"},
			want: "bbstest1",
		},
		{
			args: args{str: "test BBSMNAME test1", substr: "BBSMNAME", repl: "bbs"},
			want: "testbbstest1",
		},
		{
			args: args{str: "test BBSMNAME", substr: "BBSMNAME", repl: "bbs"},
			want: "testbbs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regexReplace(tt.args.str, tt.args.substr, tt.args.repl); got != tt.want {
				t.Errorf("regexReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}
