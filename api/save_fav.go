package api

import "github.com/PichuChen/go-bbs"

type SaveFavParams struct {
	Fav    *bbs.Fav
	UserID string
}

type SaveFavResult struct {
	MTime int64
	Fav   *bbs.Fav
}

func SaveFav(userID string, params interface{}) (interface{}, error) {
	saveFavParams, ok := params.(*SaveFavParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	if userID != saveFavParams.UserID {
		return nil, ErrInvalidParams
	}

	fav, mtime, err := saveFavParams.Fav.Save(saveFavParams.UserID)
	result := &SaveFavResult{
		MTime: mtime,
		Fav:   fav,
	}

	return result, err
}
