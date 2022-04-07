package admin

import (
	"cms/model"
	"cms/util"
	"math"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	UserName string `form:"uname" error-required:"请填写用户名"  error-strlen:"请填写用户名" validate:"required,strlen=5"`
	Password string `form:"password"`
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

	//注册自定义验证规则
	validate.RegisterValidation("strlen", util.StrLenFunc)

	if err := validate.Struct(r); err != nil {
		util.ValidateErrHandle(ctx, r, err)
		return
	}

	id := ctx.PostValueIntDefault("id", 0)
	now := uint(time.Now().Unix())

	data := model.User{
		UserName: r.UserName,
	}

	if id > 0 {
		data.ID = uint(id)
	} else {
		if len(r.Password) < 6 {
			util.Response.Fail(ctx, "请输入至少6位密码")
			return
		}

		data.AddTime = now
	}

	data.Dateline = now
	data.IsShow = 1

	//密码处理
	if r.Password != "" {
		if len(r.Password) < 6 {
			util.Response.Fail(ctx, "请输入至少6位密码")
			return
		}

		hashedPassword, err := util.EncryptedPassword(r.Password)
		if err != nil {
			util.Response.Fail(ctx, "操作失败"+err.Error())
			return
		}
		data.Password = hashedPassword
	}

	_, err := model.User{}.Save(data)
	if err != nil {
		if index := strings.Index(err.Error(), "UNIQUE"); index > -1 {
			util.Response.Fail(ctx, "用户名重复")
			return
		}

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
