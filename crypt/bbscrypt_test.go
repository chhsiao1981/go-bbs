package crypt

import (
	"reflect"
	"testing"
)

func TestFcrypt(t *testing.T) {
	type args struct {
		input    []byte
		expected []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				input:    []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', 0},
				expected: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65},
			},
			want: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fcrypt(tt.args.input, tt.args.expected)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fcrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fcrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
