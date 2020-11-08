package bbs

import (
	"reflect"
	"testing"
)

func TestFavLoad(t *testing.T) {
	origHomeDir := HOME_DIR
	HOME_DIR = "./testcase"
	defer resetHomeDir(origHomeDir)

	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *Fav
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{userID: "testUser"},
			want: &Fav{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FavLoad(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FavLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FavLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
