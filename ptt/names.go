package ptt

import "unicode"

func isValidUserID(userID string) bool {
	if len(userID) < 2 || len(userID) > PTT_IDLEN {
		return false
	}

	if !isalpha(userID[0]) {
		return false
	}

	for _, c := range userID {
		if !unicode.IsNumber(c) && !unicode.IsLetter(c) {
			return false
		}
	}

	return true
}

func isalpha(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}
