package admin

import (
	"github.com/kataras/iris/v12"
)

//默认首页
func InitIndex(party iris.Party) {
	party.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/admin/login/index")
		return
	})
}
