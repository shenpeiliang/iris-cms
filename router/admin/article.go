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
	article.Get("/lists", hero.Handler(articleController.Lists))
}
