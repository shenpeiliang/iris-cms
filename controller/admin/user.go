package admin

import (
	"bytes"
	"reflect"
	"strconv"
	"time"

	"cms/middleware"
	"cms/util"

	"github.com/dchest/captcha"
	"gopkg.in/go-playground/validator.v9"

	"github.com/kataras/iris/v12"
)

type User struct {
	UserName string `json:"admin_name" form:"admin_name" error-required:"请填写用户名" error-strlen:"请填写用户名" validate:"required,strlen=5"`
	Password string `json:"password" form:"password" error-required:"请填写用户密码" validate:"required"`
	Code     string `json:"code" form:"code" error-len:"请填写6位长度的验证码" error-required:"请填写6位长度的验证码" validate:"required,len=6"`
}

//登录
func (u User) Login(ctx iris.Context) {
	ctx.View("admin/user/login.html")
}

//自定义验证规则
func strLenFunc(fl validator.FieldLevel) bool {
	param, _ := strconv.Atoi(fl.Param())

	valueLen := len(fl.Field().String())

	if valueLen < param {
		return false
	}
	return true
}

//登录检查
func (u User) Check(ctx iris.Context) {
	//表单数据转struct
	/* if err := ctx.ReadJSON(&u); err != nil {
		ctx.JSON(map[string]string{
			"code": "fail",
			"msg":  err.Error(),
		})

		return
	} */

	data := &User{
		UserName: ctx.PostValueDefault("admin_name", ""),
		Password: ctx.PostValueDefault("password", ""),
		Code:     ctx.PostValueDefault("code", ""),
	}

	validate := validator.New()

	//注册自定义验证规则
	validate.RegisterValidation("strlen", strLenFunc)

	if err := validate.Struct(data); err != nil {
		//是否空值
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.JSON(map[string]string{
				"code": strconv.Itoa(iris.StatusInternalServerError),
				"msg":  err.Error(),
			})
			return
		}

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fieldName := e.Field()
			//反射获取其他标签信息
			field, ok := reflect.TypeOf(u).FieldByName(fieldName)
			errInfo := field.Tag.Get("error-" + e.Tag())

			if ok {
				util.Response.Fail(ctx, errInfo)
				return
			}
		}

		util.Response.Fail(ctx, err.Error())
		return
	}

	//session缓存
	id := middleware.Session.Start(ctx).GetStringDefault("captcha", "")

	code := ctx.PostValueDefault("code", "")

	if ok := captcha.VerifyString(id, code); !ok {
		util.Response.Fail(ctx, "验证码错误")
		return
	}

	result := iris.Map{
		"url": "/admin/article/lists",
	}

	util.Response.Success(ctx, "登录成功", result)
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
