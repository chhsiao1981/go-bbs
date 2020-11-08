package bbs

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Utf8ToBig5(input []rune) []byte {
	utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
	big5, _, _ := transform.String(utf8ToBig5, string(input))
	return []byte(big5)
}

func Big5ToUtf8(input []byte) []rune {
	big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
	utf8, _, _ := transform.String(big5ToUTF8, string(input))
	return []rune(utf8)
}
