package admin

import (
	"cms/controller/admin"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func RegisterLogin(party iris.Party) {
	user := party.Party("/login")
	loginController := admin.Login{}

	user.Get("/index", hero.Handler(loginController.Index)).Name = "/admin/login/index"
	user.Post("/check", hero.Handler(loginController.Check)).Name = "/admin/login/check"
	user.Get("/captcha", hero.Handler(loginController.Captcha)).Name = "/admin/login/captcha"
}
