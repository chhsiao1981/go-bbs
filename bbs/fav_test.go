package bbs

import (
	"reflect"
	"testing"
)

/*
test0:
[35, 13,
 3, 0, 2, 1,
 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 3, 1, 1,
 2, 1, 1, 183, 115, 170, 186, 165, 216, 191, 253, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
       0, 0, 0, 0, 0, 0, 0, 0, 0,
 3, 1, 2,
 1, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 1, 1, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

 1, 0, 0, 0,
 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
*/

var (
	testTitle0 = [BTLEN + 1]byte{
		183, 115, 170, 186, 165, 216, 191, 253, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0,
	}
	testSubFav0 = &Fav{
		NAllocs:  9,
		DataTail: 1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*FavType{
			&FavType{FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
	}
	testFav0 = &Fav{
		NAllocs:  14,
		DataTail: 6,
		NBoards:  3,
		NLines:   2,
		NFolders: 1,
		LineID:   2,
		FolderID: 1,
		Favh: []*FavType{
			&FavType{FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			&FavType{FAVT_LINE, 1, &FavLine{1}},
			&FavType{FAVT_FOLDER, 1, &FavFolder{1, testTitle0, testSubFav0}},
			&FavType{FAVT_LINE, 1, &FavLine{2}},
			&FavType{FAVT_BOARD, 1, &FavBoard{9, 0, 0}},
			&FavType{FAVT_BOARD, 1, &FavBoard{8, 0, 0}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
	}
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
			name: "testUser",
			args: args{userID: "testUser"},
			want: testFav0,
		},
		{
			name: "testNonExists",
			args: args{userID: "testNonExists"},
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
			if got != nil {
				for idx, each := range got.Favh {
					eachWant := tt.want.Favh[idx]
					if !reflect.DeepEqual(each, eachWant) {
						if each.TheType == FAVT_FOLDER {
							eachFolder := each.Fp.(*FavFolder)
							eachWantFolder := eachWant.Fp.(*FavFolder)
							t.Errorf("FolderError: eachFolder: %v eachWantFolder: %v", eachFolder.ThisFolder, eachWantFolder.ThisFolder)
						}
					}
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FavLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFavSave(t *testing.T) {
	origHomeDir := HOME_DIR
	HOME_DIR = "./testcase"
	defer resetHomeDir(origHomeDir)

	type args struct {
		fav    *Fav
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *Fav
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{fav: testFav0, userID: "testUserWrite"},
			want: testFav0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FavSave(tt.args.fav, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("FavSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := FavLoad(tt.args.userID)
			if err != nil {
				t.Errorf("FavSave(): unable to load: got: %v e: %v", got, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FavSave(): got: %v testFav0: %v", got, tt.want)
			}
		})
	}
}
