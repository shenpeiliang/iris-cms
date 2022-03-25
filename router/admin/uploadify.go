package admin

import (
	"cms/controller/admin"
	"cms/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func RegisterUploadify(party iris.Party) {
	uploadify := party.Party("/uploadify", middleware.Auth, middleware.Common)

	uploadifyController := admin.Uploadify{}
	uploadify.Post("/upload", hero.Handler(uploadifyController.Upload)).Name = "/admin/uploadify/upload"

}
