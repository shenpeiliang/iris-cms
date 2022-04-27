package main

import (
	"cms/router"
	"cms/util"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	//初始化配置
	config := util.Config{}.New()

	//日志模式
	leverName := "disable"
	if s, has := config.GetOther()["Server"]; has {
		item := s.(map[string]interface{})
		if v, has := item["LogLevel"]; has {
			leverName = v.(string)
		}
	}
	app.Logger().SetLevel(leverName)

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
