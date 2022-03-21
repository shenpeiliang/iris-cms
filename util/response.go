package util

import (
	"github.com/kataras/iris/v12"
)

type ResponseHandle struct {
}

var Response = new(ResponseHandle)

//返回信息
func (r ResponseHandle) Json(ctx iris.Context, code, msg string, data map[string]interface{}) {
	content := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	ctx.JSON(content)
}

//成功返回
func (r ResponseHandle) Success(ctx iris.Context, msg string, data ...map[string]interface{}) {
	result := make(map[string]interface{})
	if len(data) > 0 {
		result = data[0]
	}
	r.Json(ctx, "200", msg, result)
}

//失败返回
func (r ResponseHandle) Fail(ctx iris.Context, msg string) {
	result := make(map[string]interface{})
	r.Json(ctx, "-1", msg, result)
}
