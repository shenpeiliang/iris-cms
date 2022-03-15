package admin

import (
	"github.com/kataras/iris/v12"
)

type Article struct {
}

//列表
func (a Article) Lists(ctx iris.Context) {
	// Bind: {{.message}} with "Hello world!"
	ctx.ViewData("message", "Hello world!")
	// Render template file: ./views/hello.html
	ctx.View("admin/Index/index.html")
}
