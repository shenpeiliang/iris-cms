package service

import (
	"cms/model"

	"github.com/kataras/iris/v12/sessions"
)

var (
	Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})

	//默认登录用户
	SessionUser = model.User{}
)
