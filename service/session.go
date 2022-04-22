package service

import (
	"cms/model"

	"github.com/kataras/iris/v12/sessions"
)

var (
	//默认登录用户
	SessionUser = model.User{}
)

//初始化session
func InitSession() *sessions.Sessions {
	s := sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})

	s.UseDatabase(GetRedisDB())

	return s
}
