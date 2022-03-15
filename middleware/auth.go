package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var Session = sessions.New(sessions.Config{
	Cookie: sessions.DefaultCookieName,
})

//登录检查
func Auth(ctx iris.Context) {
	session := Session.Start(ctx)

	uid := session.GetIntDefault("uid", 0)

	ctx.Values().Set("uid", uid)

	ctx.ViewData("uid", uid)

	ctx.Next()
}
