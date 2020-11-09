package ptt

import (
	"testing"
)

func Test_isalpha(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{'/'},
			want: false,
		},
		{
			args: args{'0'},
			want: false,
		},
		{
			args: args{'9'},
			want: false,
		},
		{
			args: args{':'},
			want: false,
		},
		{
			args: args{'@'},
			want: false,
		},
		{
			args: args{'A'},
			want: true,
		},
		{
			args: args{'Z'},
			want: true,
		},
		{
			args: args{'['},
			want: false,
		},
		{
			args: args{'`'},
			want: false,
		},
		{
			args: args{'a'},
			want: true,
		},
		{
			args: args{'z'},
			want: true,
		},
		{
			args: args{'{'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isalpha(tt.args.c); got != tt.want {
				t.Errorf("isalpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidUserID(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "too short",
			args: args{"a"},
			want: false,
		},
		{
			name: "good",
			args: args{"a1"},
			want: true,
		},
		{
			name: "good",
			args: args{"a12345678901"},
			want: true,
		},
		{
			name: "too long",
			args: args{"a1234567890123"},
			want: false,
		},
		{
			name: "1st not as alpha",
			args: args{"0123456789012"},
			want: false,
		},
		{
			name: "require alnum",
			args: args{"a12345678-"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidUserID(tt.args.userID); got != tt.want {
				t.Errorf("isValidUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
