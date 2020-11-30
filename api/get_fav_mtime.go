package api

import "github.com/PichuChen/go-bbs"

type GetFavMTimeParams struct {
	UserID string
}

type GetFavMTimeResult struct {
	MTime int64
}

func GetFavMTime(userID string, params interface{}) (interface{}, error) {
	getFavParams, ok := params.(*GetFavMTimeParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	if userID != getFavParams.UserID {
		return nil, ErrInvalidParams
	}

	mtime, err := bbs.FavMTime(getFavParams.UserID)
	if err != nil {
		return nil, err
	}

	result := &GetFavMTimeResult{
		MTime: mtime,
	}

	return result, nil
}
