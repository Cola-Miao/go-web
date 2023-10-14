package pkg

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func EncodeString(s string) (es []byte, err error) {
	es, err = bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return
}

func CompareEncodeString(es string, s string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(es), []byte(s))
	if err != nil {
		err = errors.New("name or password wrong")
	}
	return
}
