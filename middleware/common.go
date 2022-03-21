package middleware

import (
	"strings"

	"github.com/kataras/iris/v12"
)

//活动菜单
type ActiveMenu struct {
	Parent string
	Child  string
}

//路由信息
type RouterData struct {
	CurrentPath    string
	ControllerPath string
}

//初始化公共数据
func Common(ctx iris.Context) {
	//初始化菜单
	ctx.ViewData("ActiveMenu", ActiveMenu{
		Parent: "nav_selected",
		Child:  " class=\"curent\" ",
	})

	//路由信息
	path := ctx.Path()
	contrllerPath := path[0 : strings.LastIndex(path, "/")+1]

	ctx.ViewData("RouterData", RouterData{
		CurrentPath:    path,
		ControllerPath: contrllerPath,
	})

}
