package bbs

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

var (
	testUserecRaw = &ptttype.UserecRaw{
		Version:    ptttype.PASSWD_VERSION,
		UserID:     [ptttype.IDLEN + 1]byte{0x53, 0x59, 0x53, 0x4f, 0x50},                          // SYSOP
		RealName:   [ptttype.REALNAMESZ]byte{0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e}, // CodingMan
		Nickname:   [ptttype.NICKNAMESZ]byte{0xaf, 0xab},                                           // 神
		PasswdHash: [ptttype.PASSLEN]byte{0x62, 0x68, 0x77, 0x76, 0x4f, 0x4a, 0x74, 0x66, 0x54, 0x31, 0x54, 0x41, 0x49, 0x00},

		UFlag:        33557088,
		UserLevel:    536871943,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     [ptttype.IPV4LEN + 1]byte{0x35, 0x39, 0x2e, 0x31, 0x32, 0x34, 0x2e, 0x31, 0x36, 0x37, 0x2e, 0x32, 0x32, 0x36},                                                     //59.124.167.226
		Address:      [ptttype.ADDRESSSZ]byte{0xb7, 0x73, 0xa6, 0xcb, 0xbf, 0xa4, 0xa4, 0x6c, 0xb5, 0xea, 0xb6, 0x6d, 0xaf, 0x51, 0xa6, 0xb3, 0xa7, 0xf8, 0x35, 0x34, 0x31, 0xb8, 0xb9}, //新竹縣子虛鄉烏有村543號
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       [ptttype.CAREERSZ]byte{0xa5, 0xfe, 0xb4, 0xba, 0xb3, 0x6e, 0xc5, 0xe9}, //全景軟體
		LastSeen:     1600681288,
	}

	testUserec1 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		Userid:       "SYSOP",
		Realname:     "CodingMan",
		Nickname:     "神",
		Passwd:       "bhwvOJtfT1TAI",
		Uflag:        33557088,
		Userlevel:    536871943,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		Lasthost:     "59.124.167.226",
	}

	testUserec2 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		Userid:       "CodingMan",
		Realname:     "朱元璋",
		Nickname:     "程式俠",
		Passwd:       "u8mLG.ktfOk3w",
		Uflag:        33557216,
		Userlevel:    31,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600737659,
		Lastlogin:    1600737960,
		Lasthost:     "59.124.167.226",
	}

	testUserec3 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		Userid:       "pichu",
		Realname:     "Pichu",
		Nickname:     "Pichu",
		Passwd:       "KO27TyME.3/tw",
		Uflag:        33557216,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600755675,
		Lastlogin:    1600766204,
		Lasthost:     "103.246.218.43",
	}

	testUserec4 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		Userid:       "Kahou",
		Realname:     "林嘉豪",
		Nickname:     "Kahou",
		Passwd:       "V3nkaYTLnDPUA",
		Uflag:        33557216,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600758266,
		Lastlogin:    1600758266,
		Lasthost:     "180.217.174.18",
	}

	testUserec5 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		Userid:       "Kahou2",
		Realname:     "Kahou",
		Nickname:     "kahou",
		Passwd:       "R7shIAOZgQCKs",
		Uflag:        33557216,
		Userlevel:    31,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600758939,
		Lastlogin:    1600760401,
		Lasthost:     "180.217.174.18",
	}
	testUserecEmpty = &Userec{}

	testOpenUserecFile1     []*Userec = nil
	TEST_N_OPEN_USER_FILE_1           = 50
)

var (
	testBoard0 = &BoardHeader{
		BrdName:    "SYSOP",
		Title:      "嘰哩 ◎站長好!",
		Brdattr:    ptttype.BRD_POSTMASK,
		Gid:        2,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard1 = &BoardHeader{
		BrdName:    "1...........",
		Title:      ".... Σ中央政府  《高壓危險,非人可敵》",
		Brdattr:    ptttype.BRD_GROUPBOARD,
		Level:      ptttype.PERM_SYSOP,
		Gid:        1,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard2 = &BoardHeader{
		BrdName:    "junk",
		Title:      "發電 ◎雜七雜八的垃圾",
		Level:      ptttype.PERM_SYSOP,
		Gid:        2,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard3 = &BoardHeader{
		BrdName:    "Security",
		Title:      "發電 ◎站內系統安全",
		Level:      ptttype.PERM_SYSOP,
		Gid:        2,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard4 = &BoardHeader{
		BrdName:    "2...........",
		Title:      ".... Σ市民廣場     報告  站長  ㄜ！",
		Brdattr:    ptttype.BRD_GROUPBOARD,
		Level:      0,
		Gid:        1,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard5 = &BoardHeader{
		BrdName:    "ALLPOST",
		Title:      "嘰哩 ◎跨板式LOCAL新文章",
		Level:      ptttype.PERM_SYSOP,
		Brdattr:    ptttype.BRD_POSTMASK,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard6 = &BoardHeader{
		BrdName:    "deleted",
		Title:      "嘰哩 ◎資源回收筒",
		Level:      ptttype.PERM_BM,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard7 = &BoardHeader{
		BrdName:    "Note",
		Title:      "嘰哩 ◎動態看板及歌曲投稿",
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard8 = &BoardHeader{
		BrdName:    "Record",
		Title:      "嘰哩 ◎我們的成果",
		Brdattr:    ptttype.BRD_POSTMASK,
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard9 = &BoardHeader{
		BrdName:    "WhoAmI",
		Title:      "嘰哩 ◎呵呵，猜猜我是誰！",
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard10 = &BoardHeader{
		BrdName:    "EditExp",
		Title:      "嘰哩 ◎範本精靈投稿區",
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard11 = &BoardHeader{
		BrdName:    "ALLHIDPOST",
		Title:      "嘰哩 ◎跨板式LOCAL新文章(隱板)",
		Brdattr:    ptttype.BRD_POSTMASK | ptttype.BRD_HIDE,
		Level:      ptttype.PERM_SYSOP,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
		BM:         make([]string, 0),
	}

	testBoard12 = &BoardHeader{
		BrdName:    "ptt_app",
		Title:      "測試 ◎大家來玩吧",
		BM:         []string{"SYSOP"},
		Brdattr:    ptttype.BRD_CPLOG | ptttype.BRD_GUESTPOST | ptttype.BRD_POSTMASK,
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
	}

	testBoard13 = &BoardHeader{
		BrdName:    "test",
		Title:      "測試 ◎測試測試測試",
		BM:         []string{"SYSOP"},
		Brdattr:    ptttype.BRD_CPLOG | ptttype.BRD_GUESTPOST | ptttype.BRD_POSTMASK,
		Level:      0,
		Gid:        5,
		BUpdate:    types.ZERO_LOCAL_TIME,
		VTime:      types.ZERO_LOCAL_TIME,
		PermReload: types.ZERO_LOCAL_TIME,
		EndGamble:  types.ZERO_LOCAL_TIME,
		SRexpire:   types.ZERO_LOCAL_TIME,
		Next:       []int32{0, 0},
		FirstChild: []int32{0, 0},
		PostType:   make([]byte, 33),
	}
)
