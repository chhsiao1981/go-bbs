package bbs

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #include "bbscrypt.h"
import "C"
import (
	"errors"
	"sync"
)

const (
	PASSWD_LEN = 13
)

var (
	ErrInvalidCrypt = errors.New("invalid crypt")
	mu              sync.Mutex
)

func Fcrypt(input []byte, expected []byte) ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	cinput := C.CBytes(input)
	defer C.free(cinput)

	cexpected := C.CBytes(expected)
	defer C.free(cexpected)

	fpasswd, err := C.fcrypt(cinput, cexpected)
	if err != nil {
		return nil, err
	}

	passwd := make([]byte, len(fpasswd))
	nCopy := copy(passwd, fpasswd)
	if nCopy != PASSWD_LEN {
		return nil, ErrInvalidCrypt
	}
	return passwd, nil
}
