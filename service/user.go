package service

import (
	"cms/model"
	"cms/util"
	"errors"
)

type User struct {
}

//用户登录
func (u User) Login(userName, password string) (model.User, error) {
	//用户信息
	userData := model.User{}.GetUser(model.User{UserName: userName})
	if userData.ID == 0 {
		return userData, errors.New("用户不存在")
	}

	ok := util.CompareHashAndPassword(userData.Password, password)
	if !ok {
		return userData, errors.New("用户密码不正确")
	}

	return userData, nil
}
