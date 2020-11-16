package ptttype

const (
	////
	// pttstruct.h
	IDLEN   = 12 /* Length of bid/uid */
	IPV4LEN = 15 /* a.b.c.d form */

	PASS_INPUT_LEN = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	PASSLEN = 14 /* Length of encrypted passwd field */
	REGLEN  = 38 /* Length of registration data */

	REALNAMESZ = 20 /* Size of real-name field */
	NICKNAMESZ = 24 /* SIze of nick-name field */
	EMAILSZ    = 50 /* Size of email field */
	ADDRESSSZ  = 50 /* Size of address field */
	CAREERSZ   = 40 /* Size of career field */
	PHONESZ    = 20 /* Size of phone field */

	PASSWD_VERSION = 4194

	BTLEN = 48 /* Length of board title */

	TTLEN = 64 /* Length of title */
	FNLEN = 28 /* Length of filename */

	USHM_SIZE = ((MAX_ACTIVE) * 41 / 40)
	/* USHM_SIZE 比 MAX_ACTIVE 大是為了防止檢查人數上限時, 又同時衝進來
	 * 會造成找 shm 空位的無窮迴圈.
	 * 又, 因 USHM 中用 hash, 空間稍大時效率較好. */

	/* MAX_BMs is dirty hardcode 4 in mbbsd/cache.c:is_BM_cache() */
	MAX_BMs = 4 /* for BMcache, 一個看板最多幾板主 */
)
