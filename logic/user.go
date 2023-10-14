package logic

import (
	"go-web/dao/mysql"
	"go-web/model"
	"go-web/pkg"
)

func LoginAuth(lf *model.LoginForm) (err error) {
	u, err := mysql.SelectUserByName(lf.Name)
	if err != nil {
		return
	}
	err = pkg.CompareEncodeString(u.Pwd, lf.Pwd)

	return
}
