package router

import (
	"cms/function"
	"cms/router/admin"

	"github.com/kataras/iris/v12"
)

//路由注册
func RegisterRouter(app *iris.Application) {

	//路由中间件
	app.Use(myMiddleware)

	//跨域处理

	//视图文件目录 每次请求时自动重载模板
	tmpl := iris.HTML("./view", ".html").Reload(true)

	//注册试图函数
	function.RegisterTemplateFun(tmpl)
	app.RegisterView(tmpl)

	//静态文件
	app.HandleDir("/static", "./static")

	//路由分组
	app.PartyFunc("/admin", func(party iris.Party) {
		//路由注册
		admin.InitAdmin(party)

	})

	//错误请求配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("error/404.html")
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.View("error/500.html")
	})

}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
