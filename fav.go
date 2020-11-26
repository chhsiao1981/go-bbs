package bbs

import (
	"github.com/PichuChen/go-bbs/fav"
	"github.com/PichuChen/go-bbs/ptttype"
)

type Fav fav.FavRaw

func NewFavFromRaw(favRaw *fav.FavRaw) *Fav {
	newFav := Fav(*favRaw)
	return &newFav
}

func (f *Fav) ToRaw() *fav.FavRaw {
	favRaw := fav.FavRaw(*f)
	return &favRaw
}

func FavMTime(userID string) (int64, error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))

	return fav.MTime(userIDRaw)
}

func FavLoad(userID string) (*Fav, error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))

	favRaw, err := fav.Load(userIDRaw)
	if err != nil {
		return nil, err
	}

	f := NewFavFromRaw(favRaw)

	return f, nil
}

func (f *Fav) Save(userID string) (*Fav, int64, error) {
	userIDRaw := &[ptttype.IDLEN + 1]byte{}
	copy(userIDRaw[:], []byte(userID))

	favRaw := f.ToRaw()

	newFavRaw, mtime, err := favRaw.Save(userIDRaw)
	var newFav *Fav
	if newFavRaw != nil {
		newFav = NewFavFromRaw(newFavRaw)
	}

	return newFav, mtime, err
}
