package admin

import (
	"cms/model"
	"math"

	"github.com/kataras/iris/v12"
)

type Article struct {
}

//列表
func (a Article) Lists(ctx iris.Context) {

	var (
		pageNum   uint
		pageCount uint
		offset    uint
	)

	offset = 0
	pageCount = 10

	data, count, _ := model.ArticleModel.Page(offset, pageCount)

	if count > 0 {
		math.Ceil(float64(count) / float64(pageCount))
	}

	ctx.ViewData("data", iris.Map{
		"act":  "热文推荐",
		"rows": data,
		"page": pageNum,
	})

	ctx.View("admin/article/lists.html")
}

func (a Article) Form(ctx iris.Context) {
	ctx.ViewData("message", "Hello world!")
	ctx.View("admin/article/form.html")
}

func (a Article) Delete(ctx iris.Context) {

}
