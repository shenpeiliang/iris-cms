package admin

import (
	"cms/model"
	"cms/util"
	"math"
	"reflect"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	UserName string `form:"uname" error-required:"请填写用户名" validate:"required"`
	Password string `form:"password" error-required:"请填写密码" validate:"required"`
}

//列表
func (a User) Lists(ctx iris.Context) {

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
		where["user_name like ?"] = title + "%"
	}

	order := "id desc"

	data, _ := model.User{}.Page(where, offset, pageCount, order)

	count, _ := model.User{}.Count(where, order)

	if count > 0 {
		pageNum = uint(math.Ceil(float64(count) / float64(pageCount)))
	}

	ctx.ViewData("data", iris.Map{
		"act":  "用户信息",
		"rows": data,
		"page": pageNum,
	})

	ctx.View("admin/user/lists.html")
}

//表单
func (a User) Form(ctx iris.Context) {
	var (
		err error
	)

	data := model.User{}

	id := ctx.Params().GetUint64Default("id", 0)

	if id > 0 {
		data, err = model.User{}.Get(uint(id))

		if err != nil {
			util.Response.Fail(ctx, "记录不存在")
			return
		}
	}

	ctx.ViewData("data", data)
	ctx.View("admin/user/form.html")
}

//保存
func (a User) Save(ctx iris.Context) {
	r := User{
		UserName: ctx.PostValueDefault("uname", ""),
		Password: ctx.PostValueDefault("password", ""),
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

	data := model.User{
		UserName: r.UserName,
		Password: r.Password,
		Dateline: now,
		IsShow:   1,
	}

	if id > 0 {
		data.ID = uint(id)
	} else {
		data.AddTime = now
	}

	ok, _ := model.User{}.Save(data)
	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//删除
func (a User) Delete(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	ok, _ := model.User{}.Delete(uint(id))

	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//显示状态
func (a User) State(ctx iris.Context) {
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

	ok, _ := model.User{}.Update(data, where)

	if !ok {
		util.Response.Fail(ctx, "操作失败")
		return
	}

	util.Response.Success(ctx, "操作成功")
}

//选择删除
func (a User) Deletes(ctx iris.Context) {
	ids := ctx.PostValues("ids[]")

	if len(ids) < 1 {
		util.Response.Fail(ctx, "请选择要操作的记录")
		return
	}

	var num uint

	for _, id := range ids {
		ok, _ := model.User{}.Delete(gconv.Uint(id))
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
