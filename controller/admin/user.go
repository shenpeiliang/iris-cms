package admin

import (
	"github.com/kataras/iris/v12"
)

type User struct {
}

//登录页面
func (u User) Index(ctx iris.Context) {
	// Bind: {{.message}} with "Hello world!"
	ctx.ViewData("message", "Hello world!")
	// Render template file: ./views/hello.html
	ctx.View("admin/user/index.html")
}

//登录
func (u User) Login(ctx iris.Context) {
	ctx.View("admin/user/login.html")
}

//登录检查
func (u User) Check(ctx iris.Context) {
	ctx.JSON(ctx.Application().ConfigurationReadOnly().GetOther())
}
