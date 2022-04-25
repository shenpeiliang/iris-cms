package admin

import (
	"bytes"
	"time"

	"cms/service"
	"cms/util"

	"github.com/dchest/captcha"
	"gopkg.in/go-playground/validator.v9"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Login struct {
	UserName string `json:"admin_name" form:"admin_name" error-required:"请填写用户名" error-strlen:"请填写用户名" validate:"required,strlen=5"`
	Password string `json:"password" form:"password" error-required:"请填写用户密码" validate:"required"`
	Code     string `json:"code" form:"code" error-len:"请填写4位长度的验证码" error-required:"请填写4位长度的验证码" validate:"required,len=4"`
}

//登录
func (u Login) Index(ctx iris.Context) {
	ctx.View("admin/login/index.html")
}

//登录检查
func (u Login) Check(ctx iris.Context) {

	data := &Login{
		UserName: ctx.PostValueDefault("admin_name", ""),
		Password: ctx.PostValueDefault("password", ""),
		Code:     ctx.PostValueDefault("code", ""),
	}

	validate := validator.New()

	//注册自定义验证规则
	validate.RegisterValidation("strlen", util.StrLenFunc)

	if err := validate.Struct(data); err != nil {
		util.ValidateErrHandle(ctx, u, err)
		return
	}

	//session缓存
	id := sessions.Get(ctx).GetStringDefault("captcha", "")

	code := ctx.PostValueDefault("code", "")

	if ok := captcha.VerifyString(id, code); !ok {
		util.Response.Fail(ctx, "验证码错误")
		return
	}

	//用户登录
	user, err := service.User{}.Login(data.UserName, data.Password)
	if err != nil {
		util.Response.Fail(ctx, err.Error())
		return
	}

	//写入session
	sessions.Get(ctx).Set("user", user)

	result := iris.Map{
		"url": "/admin/article/lists",
	}

	util.Response.Success(ctx, "登录成功", result)
}

//退出
func (u Login) Out(ctx iris.Context) {

	//写入session
	sessions.Get(ctx).Clear()

	ctx.Redirect("/admin/login/index")

}

//图形验证码
func (u Login) Captcha(ctx iris.Context) {
	width, height := 160, 45

	//图形验证码
	//captchaId := captcha.NewLen(captcha.DefaultLen)
	captchaId := captcha.NewLen(4)

	//session缓存
	sessions.Get(ctx).Set("captcha", captchaId)

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

	ctx.ServeContent(bytes.NewReader(writer.Bytes()), captchaId+".png", time.Time{})

}
