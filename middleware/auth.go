package middleware

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})

	user = iris.Map{
		"uid":   0,
		"uname": "",
	}
)

//用户信息
func User(ctx iris.Context) {
	session := Session.Start(ctx)

	data := session.Get("user")
	if data != nil {
		user = gconv.Map(data)
	}

	ctx.Values().Set("user", user)

	ctx.ViewData("user", user)

	ctx.Next()
}

//登录检查
func Auth(ctx iris.Context) {
	data := ctx.Values().Get("user")
	if data != nil {
		user = gconv.Map(data)
	}

	if user["uid"] == 0 {
		ctx.Redirect("/admin/login/index")
		return
	}

	ctx.Next()
}
