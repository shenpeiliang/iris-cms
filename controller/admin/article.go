package admin

import (
	"cms/model"
	"cms/util"
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

		where map[string]interface{}
	)

	offset = 0
	pageCount = 10

	data, _ := model.Article{}.Page(where, offset, pageCount)

	count, _ := model.Article{}.Count(where)

	if count > 0 {
		pageNum = uint(math.Ceil(float64(count) / float64(pageCount)))
	}

	ctx.ViewData("data", iris.Map{
		"act":  "热文推荐",
		"rows": data,
		"page": pageNum,
	})

	ctx.View("admin/article/lists.html")
}

func (a Article) Form(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	data, err := model.Article{}.Get(uint(id))

	if err != nil {
		util.Response.Fail(ctx, "记录不存在")
		return
	}

	ctx.ViewData("data", data)
	ctx.View("admin/article/form.html")
}

func (a Article) Delete(ctx iris.Context) {

}
