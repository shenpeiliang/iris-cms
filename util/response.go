package util

import (
	"encoding/json"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

type ResponseHandle struct {
}

var Response = new(ResponseHandle)

//返回信息
func (r ResponseHandle) Response(ctx iris.Context, code, msg string, data map[string]interface{}) {
	content := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	htmlContent, err := json.Marshal(content)
	if err != nil {
		ctx.Application().Logger().Log(golog.Default.Level, err.Error())
	}

	ctx.Negotiation().JSON(content).XML(content).HTML(string(htmlContent))

	ctx.Negotiate(nil)
}

//成功返回
func (r ResponseHandle) Success(ctx iris.Context, msg string, data ...map[string]interface{}) {
	result := make(map[string]interface{})
	if len(data) > 0 {
		result = data[0]
	}
	r.Response(ctx, "200", msg, result)
}

//失败返回
func (r ResponseHandle) Fail(ctx iris.Context, msg string) {
	result := make(map[string]interface{})
	r.Response(ctx, "-1", msg, result)
}
