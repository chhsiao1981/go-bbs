package ptttype

import (
	"reflect"
	"testing"
)

func TestLoadBoardHeaderRawsFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		expected []*BoardHeaderRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{"./testcase/board/01.BRD"},
			expected: []*BoardHeaderRaw{testBoard0, testBoard1, testBoard2, testBoard3, testBoard4, testBoard5, testBoard6, testBoard7, testBoard8, testBoard9, testBoard10, testBoard11, testBoard12, testBoard13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadBoardHeaderRawsFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBoardHeaderRawsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("LoadBoardHeaderRawsFromFile() = %v, want %v", got, tt.expected)
			}

			for idx, eachGot := range got {
				if idx >= len(tt.expected) {
					t.Errorf("BoardName not match (%v) got: %v", idx, eachGot.Brdname)
					t.Errorf("Title not match (%v) got: %v", idx, eachGot.Title)
					t.Errorf("BM not match (%v) got: %v", idx, eachGot.BM)
					t.Errorf("BrdAttr not match (%v) got: %v", idx, eachGot.BrdAttr)
					t.Errorf("VoteLimitPosts not match (%v) got: %v", idx, eachGot.VoteLimitPosts_)
					t.Errorf("VoteLimitLogins not match (%v) got: %v", idx, eachGot.VoteLimitLogins)
					t.Errorf("ChessCountry not match (%v) got: %v", idx, eachGot.ChessCountry)
					t.Errorf("BUpdate not match (%v) got: %v", idx, eachGot.BUpdate)
					t.Errorf("PostLimitPosts not match (%v) got: %v", idx, eachGot.PostLimitPosts_)
					t.Errorf("BVote not match (%v) got: %v", idx, eachGot.BVote)
					t.Errorf("VTime not match (%v) got: %v", idx, eachGot.VTime)
					t.Errorf("Level not match (%v) got: %v", idx, eachGot.Level)
					t.Errorf("PermReload not match (%v) got: %v", idx, eachGot.PermReload)
					t.Errorf("Gid not match (%v) got: %v", idx, eachGot.Gid)
					t.Errorf("Nextnot match (%v) got: %v", idx, eachGot.Next)
					t.Errorf("FirstChild not match (%v) got: %v", idx, eachGot.FirstChild)
					t.Errorf("Parent not match (%v) got: %v", idx, eachGot.Parent)
					t.Errorf("ChildCount not match (%v) got: %v", idx, eachGot.ChildCount)
					t.Errorf("Nuser not match (%v) got: %v", idx, eachGot.NUser)
					t.Errorf("PostExpire not match (%v) got: %v", idx, eachGot.PostExpire)
					t.Errorf("EndGamble not match (%v) got: %v", idx, eachGot.EndGamble)
					t.Errorf("PostType not match (%v) got: %v", idx, eachGot.PostType)
					t.Errorf("PostTypeF not match (%v) got: %v", idx, eachGot.PostTypeF)
					t.Errorf("FastRecommendPause not match (%v) got: %v", idx, eachGot.FastRecommendPause)
					t.Errorf("VoteLimitBadPost not match (%v) got: %v", idx, eachGot.VoteLimitBadpost)
					t.Errorf("PostLimitBadPost not match (%v) got: %v", idx, eachGot.PostLimitBadpost)
					t.Errorf("SRexpire not match (%v) got: %v", idx, eachGot.SRexpire)

					continue
				}

				eachExpect := tt.expected[idx]

				if !reflect.DeepEqual(eachGot, eachExpect) {
					t.Errorf("each not match: expected: %v got: %v", eachExpect, eachGot)
				}

				if eachGot.Brdname != eachExpect.Brdname {
					t.Errorf("BoardName not match (%v) expected: %v, got: %v", idx, eachExpect.Brdname, eachGot.Brdname)
				}
				if eachGot.Title != eachExpect.Title {
					t.Errorf("Title not match (%v) expected: %v, got: %v", idx, eachExpect.Title, eachGot.Title)
				}
				if eachGot.BM != eachExpect.BM {
					t.Errorf("BM not match (%v) expected: %v, got: %v", idx, eachExpect.BM, eachGot.BM)
				}
				if eachGot.BrdAttr != eachExpect.BrdAttr {
					t.Errorf("BrdAttr not match (%v) expected: %v, got: %v", idx, eachExpect.BrdAttr, eachGot.BrdAttr)
				}
				if eachGot.VoteLimitPosts_ != eachExpect.VoteLimitPosts_ {
					t.Errorf("VoteLimitPosts not match (%v) expected: %v, got: %v", idx, eachExpect.VoteLimitPosts_, eachGot.VoteLimitPosts_)
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
				if eachGot.PostLimitPosts_ != eachExpect.PostLimitPosts_ {
					t.Errorf("PostLimitPosts not match (%v) expected: %v, got: %v", idx, eachExpect.PostLimitPosts_, eachGot.PostLimitPosts_)
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
				if eachGot.NUser != eachExpect.NUser {
					t.Errorf("Nuser not match (%v) expected: %v, got: %v", idx, eachExpect.NUser, eachGot.NUser)
				}
				if eachGot.PostExpire != eachExpect.PostExpire {
					t.Errorf("PostExpire not match (%v) expected: %v, got: %v", idx, eachExpect.PostExpire, eachGot.PostExpire)
				}
				if eachGot.EndGamble != eachExpect.EndGamble {
					t.Errorf("EndGamble not match (%v) expected: %v, got: %v", idx, eachExpect.EndGamble, eachGot.EndGamble)
				}
				if eachGot.PostType != eachExpect.PostType {
					t.Errorf("PostType not match (%v) expected: %v, got: %v", idx, eachExpect.PostType, eachGot.PostType)
				}
				if eachGot.PostTypeF != eachExpect.PostTypeF {
					t.Errorf("PostTypeF not match (%v) expected: %v, got: %v", idx, eachExpect.PostTypeF, eachGot.PostTypeF)
				}
				if eachGot.FastRecommendPause != eachExpect.FastRecommendPause {
					t.Errorf("FastRecommendPause not match (%v) expected: %v, got: %v", idx, eachExpect.FastRecommendPause, eachGot.FastRecommendPause)
				}
				if eachGot.VoteLimitBadpost != eachExpect.VoteLimitBadpost {
					t.Errorf("VoteLimitBadPost not match (%v) expected: %v, got: %v", idx, eachExpect.VoteLimitBadpost, eachGot.VoteLimitBadpost)
				}
				if eachGot.PostLimitBadpost != eachExpect.PostLimitBadpost {
					t.Errorf("PostLimitBadPost not match (%v) expected: %v, got: %v", idx, eachExpect.PostLimitBadpost, eachGot.PostLimitBadpost)
				}
				if eachGot.SRexpire != eachExpect.SRexpire {
					t.Errorf("SRexpire not match (%v) expected: %v, got: %v", idx, eachExpect.SRexpire, eachGot.SRexpire)
				}
			}
		})
	}
}
