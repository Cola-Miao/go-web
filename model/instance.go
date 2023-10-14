package model

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string `validate:"min=1,max=16" gorm:"index"`
	Pwd       string `validate:"min=8,max=30"`
	Phone     string `validate:"omitempty,max=16"`
	Email     string `validate:"omitempty,email"`
	Desc      string `validate:"omitempty,max=128"`
	Group     int
	Point     int
	LastLogin time.Time
}

type Claim struct {
	Name string
	jwt.StandardClaims
}
