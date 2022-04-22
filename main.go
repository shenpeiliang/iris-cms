package main

import (
	"cms/router"
	"cms/util"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	//调试模式
	app.Logger().SetLevel("debug")

	//初始化配置
	config := iris.WithConfiguration(util.Config{}.GetAll())
	app.Configure(config)

	//路由注册
	router.RegisterRouter(app)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8080"))
}
