package mysql

import (
	"errors"
	"go-web/model"
	"go-web/pkg"
	"gorm.io/gorm"
	"time"
)

func SelectUserByID(id int) (user *model.User, err error) {
	err = mySQL.Where("id = ?", id).First(&user).Error
	return
}

func SelectUserByName(name string) (user *model.User, err error) {
	err = mySQL.Where("name = ?", name).First(&user).Error
	return
}

func UserExist(name string) (exist bool) {
	_, err := SelectUserByName(name)
	exist = !((err != nil) && errors.As(err, &gorm.ErrRecordNotFound))
	return
}

func CreateUser(sf *model.SignupForm) (err error) {
	if UserExist(sf.Name) {
		err = errors.New("userName exist")
		return
	}

	cryptPwd, err := pkg.EncodeString(sf.Pwd)
	user := &model.User{
		Name:      sf.Name,
		Pwd:       string(cryptPwd),
		LastLogin: time.Now(),
	}

	err = mySQL.Model(&model.User{}).Create(user).Error
	return
}
