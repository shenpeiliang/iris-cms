package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Common struct {
	Session *sessions.Session
}

//启动session
func (c Common) SessionStart(ctx iris.Context) {
	//session缓存
	c.Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	}).Start(ctx)
}
