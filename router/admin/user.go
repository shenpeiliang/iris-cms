package admin

import (
	"cms/controller/admin"
	"cms/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func RegisterUser(party iris.Party) {
	user := party.Party("/user", middleware.Auth, middleware.Common)
	userController := admin.User{}

	user.Get("/lists", hero.Handler(userController.Lists)).Name = "/admin/user/lists"
	user.Post("/save", hero.Handler(userController.Save)).Name = "/admin/user/save"
	user.Get("/add", hero.Handler(userController.Form)).Name = "/admin/user/add"
	user.Get("/edit/{id:uint64}", hero.Handler(userController.Form)).Name = "/admin/user/edit"
	user.Get("/delete/{id:uint64}", hero.Handler(userController.Delete)).Name = "/admin/user/delete"
}
