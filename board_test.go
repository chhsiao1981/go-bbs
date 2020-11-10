package bbs

import (
	"reflect"
	"testing"
)

/*
func TestSomething(t *testing.T) {
	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}

	expected := []BoardHeader{
		{
			BrdName: "SYSOP",
			Title:   "嘰哩 ◎站長好!",

			Brdattr: PTT_BRD_POSTMASK,
			Gid:     2,
		},
		{
			BrdName:            "1...........",
			Title:              ".... Σ中央政府  《高壓危險,非人可敵》",
			BM:                 "",
			Brdattr:            PTT_BRD_GROUPBOARD,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                1,
			Next:               []int32{0, 0},
			FirstChild:         []int32{0, 0},
			Parent:             0,
			ChildCount:         0,
			Nuser:              0,
			PostExpire:         0,
			EndGamble:          time.Unix(int64(0), 0),
			PostType:           "",
			PostTypeF:          "",
			FastRecommendPause: 0,
			VoteLimitBadPost:   0,
			PostLimitBadPost:   0,
			SRexpire:           time.Unix(int64(0), 0),
		},
		{
			BrdName:            "junk",
			Title:              "發電 ◎雜七雜八的垃圾",
			BM:                 "",
			Brdattr:            0,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                2,
			Next:               []int32{0, 0},
			FirstChild:         []int32{0, 0},
			Parent:             0,
			ChildCount:         0,
			Nuser:              0,
			PostExpire:         0,
			EndGamble:          time.Unix(int64(0), 0),
			PostType:           "",
			PostTypeF:          "",
			FastRecommendPause: 0,
			VoteLimitBadPost:   0,
			PostLimitBadPost:   0,
			SRexpire:           time.Unix(int64(0), 0),
		},
		{
			BrdName:            "Security",
			Title:              "發電 ◎站內系統安全",
			BM:                 "",
			Brdattr:            0,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                2,
			Next:               []int32{0, 0},
			FirstChild:         []int32{0, 0},
			Parent:             0,
			ChildCount:         0,
			Nuser:              0,
			PostExpire:         0,
			EndGamble:          time.Unix(int64(0), 0),
			PostType:           "",
			PostTypeF:          "",
			FastRecommendPause: 0,
			VoteLimitBadPost:   0,
			PostLimitBadPost:   0,
			SRexpire:           time.Unix(int64(0), 0),
		},
		{
			BrdName:            "2...........",
			Title:              ".... Σ市民廣場     報告  站長  ㄜ！",
			BM:                 "",
			Brdattr:            PTT_BRD_GROUPBOARD,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                2,
			Next:               []int32{0, 0},
			FirstChild:         []int32{0, 0},
			Parent:             0,
			ChildCount:         0,
			Nuser:              0,
			PostExpire:         0,
			EndGamble:          time.Unix(int64(0), 0),
			PostType:           "",
			PostTypeF:          "",
			FastRecommendPause: 0,
			VoteLimitBadPost:   0,
			PostLimitBadPost:   0,
			SRexpire:           time.Unix(int64(0), 0),
		},
	}

	for index, header := range headers[0:2] {

		if header.BrdName != expected[index].BrdName {
			t.Errorf("BoardName not match in index %d, expected: %s, got: %s", index, expected[index].BrdName, header.BrdName)
		}
		if header.Title != expected[index].Title {
			t.Errorf("Title not match in index %d, expected: %s, got: %s", index, expected[index].Title, header.Title)
		}
		if header.BM != expected[index].BM {
			t.Errorf("BM not match in index %d, expected: %s, got: %s", index, expected[index].BM, header.BM)
		}
		if header.VoteLimitPosts != expected[index].VoteLimitPosts {
			t.Errorf("VoteLimitPosts not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitPosts, header.VoteLimitPosts)
		}
		if header.VoteLimitLogins != expected[index].VoteLimitLogins {
			t.Errorf("VoteLimitLogins not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitLogins, header.VoteLimitLogins)
		}
		if header.ChessCountry != expected[index].ChessCountry {
			t.Errorf("ChessCountry not match in index %d, expected: %s, got: %s", index, expected[index].ChessCountry, header.ChessCountry)
		}
		if header.BUpdate != expected[index].BUpdate {
			t.Errorf("BUpdate not match in index %d, expected: %s, got: %s", index, expected[index].BUpdate, header.BUpdate)
		}
		if header.PostLimitPosts != expected[index].PostLimitPosts {
			t.Errorf("PostLimitPosts not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitPosts, header.PostLimitPosts)
		}
		if header.BVote != expected[index].BVote {
			t.Errorf("BVote not match in index %d, expected: %d, got: %d", index, expected[index].BVote, header.BVote)
		}
		if header.VTime != expected[index].VTime {
			t.Errorf("VTime not match in index %d, expected: %s, got: %s", index, expected[index].VTime, header.VTime)
		}
		if header.Level != expected[index].Level {
			t.Errorf("Level not match in index %d, expected: %d, got: %d", index, expected[index].Level, header.Level)
		}
		if header.PermReload != expected[index].PermReload {
			t.Errorf("PermReload not match in index %d, expected: %s, got: %s", index, expected[index].PermReload, header.PermReload)
		}
		if header.Gid != expected[index].Gid {
			t.Errorf("Gid not match in index %d, expected: %d, got: %d", index, expected[index].Gid, header.Gid)
		}
		for i := 0; i < 2; i++ {
			if header.Next[i] != expected[index].Next[i] {
				t.Errorf("Nextnot match in index %d, expected: %d, got: %d", index, expected[index].Next[i], header.Next[i])
			}
		}
		for i := 0; i < 2; i++ {
			if header.FirstChild[i] != expected[index].FirstChild[i] {
				t.Errorf("FirstChild not match in index %d, expected: %d, got: %d", index, expected[index].FirstChild[i], header.FirstChild[i])
			}
		}
		if header.Parent != expected[index].Parent {
			t.Errorf("Parent not match in index %d, expected: %d, got: %d", index, expected[index].Parent, header.Parent)
		}
		if header.ChildCount != expected[index].ChildCount {
			t.Errorf("ChildCount not match in index %d, expected: %d, got: %d", index, expected[index].ChildCount, header.ChildCount)
		}
		if header.Nuser != expected[index].Nuser {
			t.Errorf("Nuser not match in index %d, expected: %d, got: %d", index, expected[index].Nuser, header.Nuser)
		}
		if header.PostExpire != expected[index].PostExpire {
			t.Errorf("PostExpire not match in index %d, expected: %d, got: %d", index, expected[index].PostExpire, header.PostExpire)
		}
		if header.EndGamble != expected[index].EndGamble {
			t.Errorf("EndGamble not match in index %d, expected: %s, got: %s", index, expected[index].EndGamble, header.EndGamble)
		}
		if header.PostType != expected[index].PostType {
			t.Errorf("PostType not match in index %d, expected: %s, got: %s", index, expected[index].PostType, header.PostType)
		}
		if header.PostTypeF != expected[index].PostTypeF {
			t.Errorf("PostTypeF not match in index %d, expected: %s, got: %s", index, expected[index].PostTypeF, header.PostTypeF)
		}
		if header.FastRecommendPause != expected[index].FastRecommendPause {
			t.Errorf("FastRecommendPause not match in index %d, expected: %d, got: %d", index, expected[index].FastRecommendPause, header.FastRecommendPause)
		}
		if header.VoteLimitBadPost != expected[index].VoteLimitBadPost {
			t.Errorf("VoteLimitBadPost not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitBadPost, header.VoteLimitBadPost)
		}
		if header.PostLimitBadPost != expected[index].PostLimitBadPost {
			t.Errorf("PostLimitBadPost not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitBadPost, header.PostLimitBadPost)
		}
		if header.SRexpire != expected[index].SRexpire {
			t.Errorf("SRexpire not match in index %d, expected: %s, got: %s", index, expected[index].SRexpire, header.SRexpire)
		}

	}

}

*/

func TestOpenBoardHeaderFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		expected []*BoardHeader
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{"testcase/board/01.BRD"},
			expected: []*BoardHeader{testBoard0, testBoard1, testBoard2, testBoard3, testBoard4, testBoard5, testBoard6, testBoard7, testBoard8, testBoard9, testBoard10, testBoard11, testBoard12, testBoard13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenBoardHeaderFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenBoardHeaderFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("OpenBoardHeaderFile() = %v, want %v", got, tt.expected)
			}

			for idx, eachGot := range got {
				if idx >= len(tt.expected) {
					t.Errorf("BoardName not match (%v) got: %v", idx, eachGot.BrdName)
					t.Errorf("Title not match (%v) got: %v", idx, eachGot.Title)
					t.Errorf("BM not match (%v) got: %v", idx, eachGot.BM)
					t.Errorf("BrdAttr not match (%v) got: %v", idx, eachGot.Brdattr)
					t.Errorf("VoteLimitPosts not match (%v) got: %v", idx, eachGot.VoteLimitPosts)
					t.Errorf("VoteLimitLogins not match (%v) got: %v", idx, eachGot.VoteLimitLogins)
					t.Errorf("ChessCountry not match (%v) got: %v", idx, eachGot.ChessCountry)
					t.Errorf("BUpdate not match (%v) got: %v", idx, eachGot.BUpdate)
					t.Errorf("PostLimitPosts not match (%v) got: %v", idx, eachGot.PostLimitPosts)
					t.Errorf("BVote not match (%v) got: %v", idx, eachGot.BVote)
					t.Errorf("VTime not match (%v) got: %v", idx, eachGot.VTime)
					t.Errorf("Level not match (%v) got: %v", idx, eachGot.Level)
					t.Errorf("PermReload not match (%v) got: %v", idx, eachGot.PermReload)
					t.Errorf("Gid not match (%v) got: %v", idx, eachGot.Gid)
					t.Errorf("Nextnot match (%v) got: %v", idx, eachGot.Next)
					t.Errorf("FirstChild not match (%v) got: %v", idx, eachGot.FirstChild)
					t.Errorf("Parent not match (%v) got: %v", idx, eachGot.Parent)
					t.Errorf("ChildCount not match (%v) got: %v", idx, eachGot.ChildCount)
					t.Errorf("Nuser not match (%v) got: %v", idx, eachGot.Nuser)
					t.Errorf("PostExpire not match (%v) got: %v", idx, eachGot.PostExpire)
					t.Errorf("EndGamble not match (%v) got: %v", idx, eachGot.EndGamble)
					t.Errorf("PostType not match (%v) got: %v", idx, eachGot.PostType)
					t.Errorf("PostTypeF not match (%v) got: %v", idx, eachGot.PostTypeF)
					t.Errorf("FastRecommendPause not match (%v) got: %v", idx, eachGot.FastRecommendPause)
					t.Errorf("VoteLimitBadPost not match (%v) got: %v", idx, eachGot.VoteLimitBadPost)
					t.Errorf("PostLimitBadPost not match (%v) got: %v", idx, eachGot.PostLimitBadPost)
					t.Errorf("SRexpire not match (%v) got: %v", idx, eachGot.SRexpire)

					continue
				}

				eachExpect := tt.expected[idx]

				if !reflect.DeepEqual(eachGot, eachExpect) {
					t.Errorf("each not match: expected: %v got: %v", eachExpect, eachGot)
				}

				if eachGot.BrdName != eachExpect.BrdName {
					t.Errorf("BoardName not match (%v) expected: %v, got: %v", idx, eachExpect.BrdName, eachGot.BrdName)
				}
				if eachGot.Title != eachExpect.Title {
					t.Errorf("Title not match (%v) expected: %v, got: %v", idx, eachExpect.Title, eachGot.Title)
				}
				if !reflect.DeepEqual(eachGot.BM, eachExpect.BM) {
					t.Errorf("BM not match (%v) expected: %v, got: %v", idx, eachExpect.BM, eachGot.BM)
				}
				if eachGot.Brdattr != eachExpect.Brdattr {
					t.Errorf("BrdAttr not match (%v) expected: %v, got: %v", idx, eachExpect.Brdattr, eachGot.Brdattr)
				}
				if eachGot.VoteLimitPosts != eachExpect.VoteLimitPosts {
					t.Errorf("VoteLimitPosts not match (%v) expected: %v, got: %v", idx, eachExpect.VoteLimitPosts, eachGot.VoteLimitPosts)
				}
				if eachGot.VoteLimitLogins != eachExpect.VoteLimitLogins {
					t.Errorf("VoteLimitLogins not match (%v) expected: %v, got: %v", idx, eachExpect.VoteLimitLogins, eachGot.VoteLimitLogins)
				}
				if eachGot.ChessCountry != eachExpect.ChessCountry {
					t.Errorf("ChessCountry not match (%v) expected: %v, got: %v", idx, eachExpect.ChessCountry, eachGot.ChessCountry)
				}
				if eachGot.BUpdate != eachExpect.BUpdate {
					t.Errorf("BUpdate not match (%v) expected: %v, got: %v", idx, eachExpect.BUpdate, eachGot.BUpdate)
				}
				if eachGot.PostLimitPosts != eachExpect.PostLimitPosts {
					t.Errorf("PostLimitPosts not match (%v) expected: %v, got: %v", idx, eachExpect.PostLimitPosts, eachGot.PostLimitPosts)
				}
				if eachGot.BVote != eachExpect.BVote {
					t.Errorf("BVote not match (%v) expected: %v, got: %v", idx, eachExpect.BVote, eachGot.BVote)
				}
				if eachGot.VTime != eachExpect.VTime {
					t.Errorf("VTime not match (%v) expected: %v, got: %v", idx, eachExpect.VTime, eachGot.VTime)
				}
				if eachGot.Level != eachExpect.Level {
					t.Errorf("Level not match (%v) expected: %v, got: %v", idx, eachExpect.Level, eachGot.Level)
				}
				if eachGot.PermReload != eachExpect.PermReload {
					t.Errorf("PermReload not match (%v) expected: %v, got: %v", idx, eachExpect.PermReload, eachGot.PermReload)
				}
				if eachGot.Gid != eachExpect.Gid {
					t.Errorf("Gid not match (%v) expected: %v, got: %v", idx, eachExpect.Gid, eachGot.Gid)
				}
				if !reflect.DeepEqual(eachGot.Next, eachExpect.Next) {
					t.Errorf("Nextnot match (%v) expected: %v, got: %v", idx, eachExpect.Next, eachGot.Next)
				}
				if !reflect.DeepEqual(eachGot.FirstChild, eachExpect.FirstChild) {
					t.Errorf("FirstChild not match (%v) expected: %v, got: %v", idx, eachExpect.FirstChild, eachGot.FirstChild)
				}
				if eachGot.Parent != eachExpect.Parent {
					t.Errorf("Parent not match (%v) expected: %v, got: %v", idx, eachExpect.Parent, eachGot.Parent)
				}
				if eachGot.ChildCount != eachExpect.ChildCount {
					t.Errorf("ChildCount not match (%v) expected: %v, got: %v", idx, eachExpect.ChildCount, eachGot.ChildCount)
				}
				if eachGot.Nuser != eachExpect.Nuser {
					t.Errorf("Nuser not match (%v) expected: %v, got: %v", idx, eachExpect.Nuser, eachGot.Nuser)
				}
				if eachGot.PostExpire != eachExpect.PostExpire {
					t.Errorf("PostExpire not match (%v) expected: %v, got: %v", idx, eachExpect.PostExpire, eachGot.PostExpire)
				}
				if eachGot.EndGamble != eachExpect.EndGamble {
					t.Errorf("EndGamble not match (%v) expected: %v, got: %v", idx, eachExpect.EndGamble, eachGot.EndGamble)
				}
				if !reflect.DeepEqual(eachGot.PostType, eachExpect.PostType) {
					t.Errorf("PostType not match (%v) expected: %v, got: %v", idx, eachExpect.PostType, eachGot.PostType)
				}
				if eachGot.PostTypeF != eachExpect.PostTypeF {
					t.Errorf("PostTypeF not match (%v) expected: %v, got: %v", idx, eachExpect.PostTypeF, eachGot.PostTypeF)
				}
				if eachGot.FastRecommendPause != eachExpect.FastRecommendPause {
					t.Errorf("FastRecommendPause not match (%v) expected: %v, got: %v", idx, eachExpect.FastRecommendPause, eachGot.FastRecommendPause)
				}
				if eachGot.VoteLimitBadPost != eachExpect.VoteLimitBadPost {
					t.Errorf("VoteLimitBadPost not match (%v) expected: %v, got: %v", idx, eachExpect.VoteLimitBadPost, eachGot.VoteLimitBadPost)
				}
				if eachGot.PostLimitBadPost != eachExpect.PostLimitBadPost {
					t.Errorf("PostLimitBadPost not match (%v) expected: %v, got: %v", idx, eachExpect.PostLimitBadPost, eachGot.PostLimitBadPost)
				}
				if eachGot.SRexpire != eachExpect.SRexpire {
					t.Errorf("SRexpire not match (%v) expected: %v, got: %v", idx, eachExpect.SRexpire, eachGot.SRexpire)
				}
			}
		})
	}
}
