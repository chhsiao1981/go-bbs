package main

import (
	"github.com/PichuChen/go-bbs/api"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Api struct {
	Func   api.ApiFunc
	Params interface{}
}

func NewApi(f api.ApiFunc, params interface{}) *Api {
	return &Api{Func: f, Params: params}
}

func (api *Api) Json(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, err)
	}

	result, err := api.Func(api.Params)
	processResult(c, result, err)
}

func (api *Api) LoginRequiredJson(c *gin.Context) {
	loginParams := &LoginRequiredParams{Data: api.Params}
	err := c.ShouldBindJSON(loginParams)
	if err != nil {
		processResult(c, nil, err)
	}

	err = verifyJwt(loginParams.UserID, loginParams.Jwt)
	if err != nil {
		processResult(c, nil, err)
	}

	result, err := api.Func(loginParams.Data)
	processResult(c, result, err)
}

func verifyJwt(userID string, raw string) error {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return ErrInvalidToken
	}

	cl := &api.JwtClaim{}
	if err := tok.Claims(api.JWT_SECRET, cl); err != nil {
		return ErrInvalidToken
	}

	if cl.UserID != userID {
		return ErrInvalidToken
	}

	return nil
}

func processResult(c *gin.Context, result interface{}, err error) {
	if err == ErrInvalidToken {
		c.JSON(401, &errResult{err.Error()})
	}
	if err != nil {
		c.JSON(500, &errResult{err.Error()})
	}

	c.JSON(200, result)
}
