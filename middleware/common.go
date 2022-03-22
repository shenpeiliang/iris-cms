package middleware

import (
	"html/template"
	"strings"

	"github.com/kataras/iris/v12"
)

type CommonData struct {
	//路由信息
	CurrentPath    string
	ControllerPath string

	//活跃菜单
	ParentMenu string
	ChildMenu  template.HTMLAttr
}

//初始化公共数据
func Common(ctx iris.Context) {
	//路由信息
	path := ctx.Path()
	contrllerPath := path[0 : strings.LastIndex(path, "/")+1]

	ctx.ViewData("CommonData", CommonData{
		CurrentPath:    path,
		ControllerPath: contrllerPath,

		ParentMenu: "nav_selected",
		ChildMenu:  template.HTMLAttr(" class='curent' "),
	})

	ctx.Next()
}
