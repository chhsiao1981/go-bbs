package fav

import "errors"

var (
	ErrInvalidFavBoard   = errors.New("invalid fav-board")
	ErrInvalidFavLine    = errors.New("invalid fav-line")
	ErrInvalidFavFolder  = errors.New("invalid fav-folder")
	ErrInvalidFavType    = errors.New("invalid fav-type")
	ErrInvalidFavRecord  = errors.New("invalid fav-record")
	ErrInvalidFav4Record = errors.New("invalid fav4-record")
	ErrOutdatedFav       = errors.New("outdated fav")
)
