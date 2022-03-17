package admin

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"cms/middleware"

	"github.com/dchest/captcha"
	"gopkg.in/go-playground/validator.v9"

	"github.com/kataras/iris/v12"
)

type User struct {
	UserName string `json:"admin_name" form:"admin_name" label:"用户名" validate:"required"`
	Password string `json:"password" form:"password" label:"密码" validate:"required"`
	Code     string `json:"code" form:"code" label:"验证码" validate:"required,len=6"`
}

//登录
func (u User) Login(ctx iris.Context) {
	ctx.View("admin/user/login.html")
}

//自定义验证规则
func UserStructValidation(sl validator.StructLevel) {
	/* user := sl.Current().Interface().(User)

	//自定义规则
	codeLen := len(user.Code)
	if codeLen != 6 {
		sl.ReportError(user.Code, "code", "Code", "len", "6")
	} */
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

	//注册验证方法
	validate.RegisterStructValidation(UserStructValidation, data)

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
			if e.StructField() == "Code" && e.Tag() == "len" {

				maxLen, _ := strconv.Atoi(e.Param())

				ctx.JSON(map[string]string{
					"code": "fail",
					"msg":  fmt.Sprintf("请填写%d位长度的验证码", maxLen),
				})
				return
			}
		}

		ctx.JSON(map[string]string{
			"code": "fail",
			"msg":  err.Error(),
		})
		return
	}

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
