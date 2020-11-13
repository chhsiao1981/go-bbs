package api

import (
	"time"

	"github.com/PichuChen/go-bbs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type LoginParams struct {
	UserID string
	Passwd string
	IP     string
}

type LoginResult struct {
	Jwt string
}

type JwtClaim struct {
	UserID string
	Expire *jwt.NumericDate
}

func Login(params interface{}) (interface{}, error) {
	loginParams, ok := params.(*LoginParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	user, err := bbs.Login(loginParams.UserID, loginParams.Passwd, loginParams.IP)
	log.Infof("after Login: user: %v e: %v", user, err)
	if err != nil {
		return nil, err
	}

	token, err := createToken(user)
	if err != nil {
		return nil, err
	}

	result := &LoginResult{
		Jwt: token,
	}

	return result, nil
}

func createToken(userec *bbs.Userec) (string, error) {
	var err error

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := &JwtClaim{
		UserID: userec.Userid,
		Expire: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}
