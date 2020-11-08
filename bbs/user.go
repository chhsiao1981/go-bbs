package bbs

import (
	"os"
	"strings"
)

func getUserHomeDir(userID string) string {
	return strings.Join([]string{HOME_DIR, "home", string(userID[0]), userID}, string(os.PathSeparator))
}

func setuserfile(userID string, postfix string) string {
	homeDir := getUserHomeDir(userID)

	return strings.Join([]string{homeDir, postfix}, string(os.PathSeparator))
}
