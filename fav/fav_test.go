package fav

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func TestFavLoad(t *testing.T) {
	setupTest()
	defer teardownTest()

	types.CopyFile("./testcase/home1", "./testcase/home")
	defer os.RemoveAll("./testcase/home")

	userID0 := [ptttype.IDLEN + 1]byte{}
	copy(userID0[:], []byte("testUser"))

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID1[:], []byte("testNoExist"))

	type args struct {
		userID *[ptttype.IDLEN + 1]byte
	}
	tests := []struct {
		name     string
		args     args
		expected *FavRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "testUser",
			args:     args{userID: &userID0},
			expected: testFav0,
		},
		{
			name:     "testNoExist",
			args:     args{userID: &userID1},
			expected: &FavRaw{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FavLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				tt.expected.LoadTime = got.LoadTime
				for idx, each := range got.Favh {
					eachWant := tt.expected.Favh[idx]
					if !reflect.DeepEqual(each, eachWant) {
						if each.TheType == FAVT_FOLDER {
							eachFolder := each.Fp.(*FavFolder)
							eachWantFolder := eachWant.Fp.(*FavFolder)
							t.Errorf("FolderError: eachFolder: %v eachWantFolder: %v", eachFolder.ThisFolder, eachWantFolder.ThisFolder)
						}
					}
				}
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FavLoad() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFavSave(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = types.CopyFile("./testcase/home1", "./testcase/home")
	defer os.RemoveAll("./testcase/home")

	time.Sleep(11 * time.Second)

	userID0 := [ptttype.IDLEN + 1]byte{}
	copy(userID0[:], []byte("testUserWrt"))

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID0[:], []byte("testUserWr1"))

	type args struct {
		fav    *FavRaw
		userID *[ptttype.IDLEN + 1]byte
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		sleep    time.Duration
		expected *FavRaw
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{fav: testFav0, userID: &userID0},
			expected: testFav0,
		},
		{
			name:     "test0",
			args:     args{fav: testFav1, userID: &userID1},
			expected: testFav0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(tt.sleep)
			tt.args.fav.LoadTime = types.GetCurrentMilliTS()

			got, _, err := tt.args.fav.Save(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FavSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got == nil {
				return
			}

			tt.expected.LoadTime = got.LoadTime
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FavSave(): got: %v testFav0: %v", got, tt.expected)
			}
		})
	}
}
