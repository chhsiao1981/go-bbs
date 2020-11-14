// +build default

package ptttype

const (
	// type struct requires const.
	// use +build to setup different config
	MAX_USERS = 150000 /* 最高註冊人數 */

	MAX_ACTIVE = 1024 /* 最多同時上站人數 */

	MAX_BOARD = 8192 /* 最大開板個數 */

	MAX_FRIEND = 256 /* 載入 cache 之最多朋友數目 */

	MAX_REJECT = 32 /* 載入 cache 之最多壞人數目 */

	MAX_MSGS = 10 /* 水球(熱訊)忍耐上限 */

	MAX_ADBANNER = 500 /* 最多動態看板數 */

	HOTBOARDCACHE = 0 /* 熱門看板快取 */

	MAX_FROM = 300 /* 最多故鄉數 */

	MAX_REVIEW = 7 /* 最多水球回顧 */

	NUMVIEWFILE = 14 /* 進站畫面最多數 */
)
