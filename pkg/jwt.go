package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-web/model"
	"time"
)

var (
	secret = []byte("Fffire")
	term   = time.Hour * 24 * 3
)

func GenToken(lf *model.LoginForm) (token string, err error) {
	now := time.Now()
	expires := now.Add(term)
	claim := model.Claim{
		Name: lf.Name,
		StandardClaims: jwt.StandardClaims{
			Audience:  lf.Name,
			ExpiresAt: expires.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "go-web",
			Subject:   "auth",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = t.SignedString(secret)

	return
}

func ValidateToken(token string) (claim *model.Claim, err error) {
	var tc *jwt.Token
	if tc, err = jwt.ParseWithClaims(token, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, err
	}); err != nil {
		return
	}
	if !tc.Valid {
		err = errors.New("invalid token")
		return
	}

	var ok bool
	if claim, ok = tc.Claims.(*model.Claim); !ok {
		err = errors.New("parse token failed")
		return
	}

	return
}
