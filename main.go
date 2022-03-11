package main

import (
	"cms/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	//调试模式
	app.Logger().SetLevel("debug")

	//配置初始化
	config := iris.WithConfiguration(iris.YAML("./config/config.yml"))

	//路由注册
	router.RegisterRouter(app)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8080"), config)
}
