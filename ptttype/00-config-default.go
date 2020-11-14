// +build !custom

package ptttype

const (
	// These are special config variables requiring to be const.
	// link 00-config.conf to ptttype/ to setup customized config
	//
	// type struct requires const.
	MAX_USERS = 100 /* 最高註冊人數 */

	MAX_ACTIVE = 31 /* 最多同時上站人數 */

	MAX_BOARD = 1024 /* 最大開板個數 */

	MAX_FRIEND = 256 /* 載入 cache 之最多朋友數目 */

	MAX_REJECT = 32 /* 載入 cache 之最多壞人數目 */

	MAX_MSGS = 10 /* 水球(熱訊)忍耐上限 */

	MAX_ADBANNER = 500 /* 最多動態看板數 */

	HOTBOARDCACHE = 0 /* 熱門看板快取 */

	MAX_FROM = 300 /* 最多故鄉數 */

	MAX_REVIEW = 7 /* 最多水球回顧 */

	NUMVIEWFILE = 14 /* 進站畫面最多數 */

	MAX_ADBANNER_SECTION = 10 /* 最多動態看板類別 */

	MAX_ADBANNER_HEIGHT = 11 /* 最大動態看板內容高度 */
)
