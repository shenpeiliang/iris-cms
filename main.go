package main

import (
	"cms/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	//调试模式
	app.Logger().SetLevel("debug")

	//初始化配置
	config := iris.YAML("./config/config.yml")

	//路由注册
	router.RegisterRouter(app)

	//默认地址  http://localhost:8080
	addr := ":8080"

	if s, has := config.GetOther()["Server"]; has {
		item := s.(map[string]interface{})
		if v, has := item["Address"]; has {
			addr = v.(string)
		}
	}

	app.Run(iris.Addr(addr), iris.WithConfiguration(config))
}
