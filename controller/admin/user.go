package admin

import (
	"bytes"
	"time"

	"cms/middleware"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

type User struct {
}

//登录
func (u User) Login(ctx iris.Context) {
	ctx.View("admin/user/login.html")
}

//登录检查
func (u User) Check(ctx iris.Context) {
	//session缓存

	id := middleware.Session.Start(ctx).GetStringDefault("captcha", "no data")

	code := ctx.PostValueDefault("code", "")

	if captcha.VerifyString(id, code) {
		ctx.JSON(map[string]string{
			"code": "ok",
			"msg":  "验证成功",
		})
	} else {
		ctx.JSON(map[string]string{
			"code": "fail",
			"msg":  "验证失败",
		})
	}

	//ctx.JSON(ctx.Application().ConfigurationReadOnly().GetOther())
}

//图形验证码
func (u User) Captcha(ctx iris.Context) {
	width, height := 160, 45

	//图形验证码
	captchaId := captcha.NewLen(captcha.DefaultLen)

	//session缓存
	middleware.Session.Start(ctx).Set("captcha", captchaId)

	//图形输出
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	ctx.Header("Content-Type", "image/png")

	var writer bytes.Buffer

	err := captcha.WriteImage(&writer, captchaId, width, height)
	if err != nil {
		ctx.JSON(map[string]string{
			"code": "error",
			"msg":  err.Error(),
		})
	}

	ctx.ServeContent(bytes.NewReader(writer.Bytes()), captchaId+".png", time.Time{}, true)

}
