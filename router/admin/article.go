package admin

import (
	"cms/controller/admin"
	"cms/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func RegisterArticle(party iris.Party) {
	article := party.Party("/article", middleware.Auth, middleware.Common)

	articleController := admin.Article{}
	article.Get("/lists", hero.Handler(articleController.Lists)).Name = "/admin/article/lists"
	article.Post("/save", hero.Handler(articleController.Save)).Name = "/admin/article/save"
	article.Get("/add", hero.Handler(articleController.Form)).Name = "/admin/article/add"
	article.Get("/edit/{id:uint64}", hero.Handler(articleController.Form)).Name = "/admin/article/edit"
	article.Get("/delete/{id:uint64}", hero.Handler(articleController.Delete)).Name = "/admin/article/delete"

	article.Post("/state", hero.Handler(articleController.State)).Name = "/admin/article/state"
	article.Post("/order", hero.Handler(articleController.Order)).Name = "/admin/article/order"
	article.Post("/deletes", hero.Handler(articleController.Deletes)).Name = "/admin/article/deletes"
}
