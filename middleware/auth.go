package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})
)

//用户信息
func User(ctx iris.Context) {
	session := Session.Start(ctx)

	uid := session.GetIntDefault("uid", 0)

	ctx.Values().Set("uid", uid)

	ctx.ViewData("uid", uid)

	ctx.Next()
}

//登录检查
func Auth(ctx iris.Context) {
	uid := ctx.Values().GetIntDefault("uid", 0)

	if uid == 0 {
		ctx.Redirect("/admin/login/index")
		return
	}

	ctx.Next()
}
