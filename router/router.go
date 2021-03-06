package router

import (
	"cms/function"
	"cms/middleware"
	"cms/router/admin"
	"cms/service"

	"github.com/kataras/iris/v12"
)

//路由注册
func RegisterRouter(app *iris.Application) {
	//视图文件目录 每次请求时自动重载模板
	tmpl := iris.HTML("./view", ".html").Reload(true)

	//注册试图函数
	function.RegisterTemplateFun(tmpl)
	app.RegisterView(tmpl)

	//静态文件
	app.HandleDir("/static", "./static")

	//路由分组
	app.PartyFunc("/", func(party iris.Party) {
		//session初始化 请求记录
		party.Use(service.InitSession().Handler(), middleware.LogRequest)

		//路由注册
		admin.InitIndex(party) //首页
		admin.InitAdmin(party) //后台管理

	})

	//错误请求配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("error/404.html")
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.View("error/500.html")
	})

}
