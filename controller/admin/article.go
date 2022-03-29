package admin

import (
	"cms/model"
	"cms/util"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

type Article struct {
	Title       string `form:"title" error-required:"请填写用户名" validate:"required"`
	Description string `form:"description"`
	Content     string `form:"content" error-required:"请填写内容" validate:"required"`
	Img         string `form:"img"`
}

//列表
func (a Article) Lists(ctx iris.Context) {

	var (
		pageNum   uint
		pageCount uint
		offset    uint
		title     string
		where     = make(map[string]interface{})
	)

	offset = 0
	pageCount = 10

	title = ctx.URLParamDefault("keyword", "")
	if title != "" {
		where["title like ?"] = title + "%"
	}

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

//表单
func (a Article) Form(ctx iris.Context) {
	var (
		err error
	)

	data := model.Article{}

	id := ctx.Params().GetUint64Default("id", 0)

	if id > 0 {
		data, err = model.Article{}.Get(uint(id))

		if err != nil {
			util.Response.Fail(ctx, "记录不存在")
			return
		}
	}

	ctx.ViewData("data", data)
	ctx.View("admin/article/form.html")
}

//保存
func (a Article) Save(ctx iris.Context) {
	r := Article{
		Title:       ctx.PostValueDefault("title", ""),
		Img:         ctx.PostValueDefault("img", ""),
		Description: ctx.PostValueDefault("description", ""),
		Content:     ctx.PostValueDefault("content", ""),
	}

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		//是否空值
		if _, ok := err.(*validator.InvalidValidationError); ok {
			util.Response.Fail(ctx, err.Error())

			return
		}

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fieldName := e.Field()
			//反射获取其他标签信息
			field, ok := reflect.TypeOf(a).FieldByName(fieldName)
			errInfo := field.Tag.Get("error-" + e.Tag())

			if ok {
				util.Response.Fail(ctx, errInfo)
				return
			}
		}

		util.Response.Fail(ctx, err.Error())
		return
	}

	id := ctx.PostValueIntDefault("id", 0)
	now := uint(time.Now().Unix())

	data := model.Article{
		Title:       r.Title,
		Content:     r.Content,
		Img:         r.Img,
		Description: r.Description,
		Dateline:    now,
		IsShow:      1,
	}

	if data.Description != "" {
		data.Description = gstr.SubStrRune(data.Description, 0, 150)
	} else {
		data.Description = util.String{}.SubHtmlText(data.Content, 0, 150)
	}

	if id > 0 {
		data.ID = uint(id)
	} else {
		data.AddTime = now
	}

	ok, _ := model.Article{}.Save(data)
	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//删除
func (a Article) Delete(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	ok, _ := model.Article{}.Delete(uint(id))

	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//显示状态
func (a Article) State(ctx iris.Context) {
	id := ctx.PostValueInt64Default("id", 0)
	if id < 1 {
		util.Response.Fail(ctx, "参数错误")
		return
	}

	state := ctx.PostValueIntDefault("is_show", 0)

	data := iris.Map{
		"is_show": byte(state),
	}

	where := iris.Map{
		"id": uint(id),
	}

	ok, _ := model.Article{}.Update(data, where)

	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//排序
func (a Article) Order(ctx iris.Context) {
	items := ctx.PostValues("arr")
	fmt.Println(items)
	if len(items) < 1 {
		util.Response.Fail(ctx, "参数错误")
		return
	}

	var num uint

	for id, paixu := range items {
		data := iris.Map{
			"paixu": paixu,
		}

		where := iris.Map{
			"id >": id,
		}

		ok, _ := model.Article{}.Update(data, where)
		if ok {
			num++
		}
	}

	if num < 1 {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//选择删除
func (a Article) Deletes(ctx iris.Context) {
	id := ctx.PostValueInt64Default("id", 0)
	if id < 1 {
		util.Response.Fail(ctx, "参数错误")
		return
	}

	state := ctx.PostValueIntDefault("is_show", 0)

	data := model.Article{
		ID:     uint(id),
		IsShow: byte(state),
	}

	ok, _ := model.Article{}.Save(data)

	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}
