package admin

import (
	"cms/controller/admin"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func UserController(party iris.Party) {
	user := party.Party("/user")
	userController := admin.User{}

	user.Get("/login", hero.Handler(userController.Login))
	user.Get("/index", hero.Handler(userController.Index))
	user.Get("/check", hero.Handler(userController.Check))
	user.Get("/captcha", hero.Handler(userController.Captcha))
}
