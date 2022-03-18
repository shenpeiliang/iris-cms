package model

import (
	"cms/util"
	"errors"
)

type Member struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"column:user_name"`
	Password string
	Salt     string
	Dateline uint
}

func (m Member) Login(uname, password string) (Member, error) {
	var data Member
	util.DB.Where(&Member{
		UserName: uname,
		Password: password,
	}).First(&data)

	//记录是否存在
	if data.ID < 1 {
		return data, errors.New("用户不存在或密码错误")
	}

	//密码校验

	return data, nil
}
