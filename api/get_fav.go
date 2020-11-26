package api

import "github.com/PichuChen/go-bbs"

type GetFavParams struct {
	UserID string
}

type GetFavResult struct {
	Fav *bbs.Fav
}

func GetFav(userID string, params interface{}) (interface{}, error) {
	getFavParams, ok := params.(*GetFavParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	if userID != getFavParams.UserID {
		return nil, ErrInvalidParams
	}

	f, err := bbs.FavLoad(getFavParams.UserID)
	if err != nil {
		return nil, err
	}

	result := &GetFavResult{
		Fav: f,
	}

	return result, nil
}
