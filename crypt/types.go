package crypt

import "unsafe"

type desCBlock [8]uint8

const DES_KEY_SZ = unsafe.Sizeof(desCBlock{})

type desKeySchedule [32]uint32 // 16 * 8 = 128 ([16]desCBlock)
