package cmsys

import "github.com/PichuChen/go-bbs/ptttype"

func StringHashWithHashBits(theBytes []byte) uint32 {
	return StringHash(theBytes) % (1 << ptttype.HASH_BITS)
}

func StringHash(theBytes []byte) uint32 {
	return fnv1a32StrCase(theBytes, FNV1_32_INIT)
}

func toupper(theByte byte) byte {
	if theByte >= 'a' && theByte <= 'z' {
		return theByte - 32
	}
	return theByte
}
