package middleware

import (
	"strings"

	"github.com/kataras/iris/v12"
)

//初始化公共数据
func Common(ctx iris.Context) {
	//路由信息
	path := ctx.Path()
	contrllerPath := path[0 : strings.LastIndex(path, "/")+1]

	ctx.ViewData("CommonData", map[string]string{
		//路由信息
		"currentPath":    path,
		"controllerPath": contrllerPath,

		//活跃菜单
		"parentMenu": "nav_selected",
		"childMenu":  " class=\"curent\" ",
	})

	ctx.Next()
}
