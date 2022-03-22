package admin

import (
	"cms/controller/admin"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func RegisterUser(party iris.Party) {
	user := party.Party("/user")
	userController := admin.User{}

	user.Get("/login", hero.Handler(userController.Login)).Name = "/admin/user/login"
	user.Post("/check", hero.Handler(userController.Check)).Name = "/admin/user/check"
	user.Get("/captcha", hero.Handler(userController.Captcha)).Name = "/admin/user/captcha"
}
