package middleware

import (
	"cms/model"
	"cms/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

//用户信息
func User(ctx iris.Context) {
	//默认值
	user := service.SessionUser

	data := sessions.Get(ctx).Get("user")
	if data != nil {
		user = data.(model.User)
	}

	ctx.Values().Set("user", user)

	ctx.ViewData("user", user)

	ctx.Next()
}

//登录检查
func Auth(ctx iris.Context) {
	data := ctx.Values().Get("user")

	user := data.(model.User)

	if user.ID == 0 {
		ctx.Redirect("/admin/login/index")
		return
	}

	ctx.Next()
}
