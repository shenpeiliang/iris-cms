package admin

import (
	"github.com/kataras/iris/v12"
)

type Article struct {
}

//列表
func (a Article) Lists(ctx iris.Context) {

	ctx.ViewData("message", "Hello world!")
	ctx.View("admin/article/lists.html")
}

func (a Article) Form(ctx iris.Context) {
	ctx.ViewData("message", "Hello world!")
	ctx.View("admin/article/form.html")
}

func (a Article) Delete(ctx iris.Context) {

}
