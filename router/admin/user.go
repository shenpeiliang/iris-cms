package admin

import (
	"cms/controller/admin"
	"cms/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func UserController(party iris.Party) {
	user := party.Party("/user")
	userController := admin.User{}

	user.Get("/login", hero.Handler(userController.Login))
	user.Post("/check", hero.Handler(userController.Check))
	user.Get("/captcha", hero.Handler(userController.Captcha))

	article := party.Party("/article", middleware.Auth)
	articleController := admin.Article{}
	article.Get("/lists", hero.Handler(articleController.Lists))
}
