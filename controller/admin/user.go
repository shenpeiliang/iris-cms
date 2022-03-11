package admin

import (
	"bytes"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

type User struct {
	Common
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

//图形验证码
func (u User) Captcha(ctx iris.Context) {
	width, height := 160, 45

	//图形验证码
	captchaId := captcha.NewLen(captcha.DefaultLen)

	u.Common.SessionStart(ctx)
	u.Common.Session.Set("captcha", captchaId)

	//图形输出
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	ctx.Header("Content-Type", "image/png")

	var writer bytes.Buffer

	captcha.WriteImage(&writer, captchaId, width, height)
}
