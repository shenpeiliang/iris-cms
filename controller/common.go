package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var Session *sessions.Session

type Common struct {
}

//启动session
func SessionStart(ctx iris.Context) {
	//session缓存
	Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	}).Start(ctx)
}
