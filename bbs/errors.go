package bbs

import "errors"

var (
	ErrInvalidFavRecord  = errors.New("invalid fav record")
	ErrInvalidFavType    = errors.New("invalid fav type")
	ErrInvalidFavFolder  = errors.New("invalid fav folder")
	ErrInvalidFavBoard   = errors.New("invalid fav board")
	ErrInvalidFavLine    = errors.New("invalid fav line")
	ErrInvalidFav4Record = errors.New("invalid fav4 record")
	ErrInvalidFile       = errors.New("invalid file")
)
