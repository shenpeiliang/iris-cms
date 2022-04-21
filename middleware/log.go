package middleware

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
)

//请求记录
func LogRequest(ctx iris.Context) {
	method := ctx.Method()
	uri := ctx.FullRequestURI()
	header := ctx.Request().Header

	curl := "curl"

	//请求数据
	body := make(map[string][]string)

	if method == "GET" {
		curl = curl + " -G "
		body = ctx.Request().URL.Query()
	} else {
		body = ctx.FormValues()
	}

	bodyLen := len(body)
	if bodyLen > 0 {
		data := make([]string, 0)
		for k, v := range body {
			data = append(data, k+"="+v[0])
		}
		curl = curl + fmt.Sprintf(" -d \"%s\" ", strings.Join(data, "&"))
	}

	//header信息
	headerLen := len(header)
	if headerLen > 0 {
		headerArr := make([]string, 0)
		for k, h := range header {
			headerArr = append(headerArr, fmt.Sprintf(" -H '%s:%s' ", k, h[0]))
		}

		curl = curl + strings.Join(headerArr, "")
	}

	curl = curl + fmt.Sprintf(" '%s' --compressed", uri)

	ctx.Application().Logger().Infof(curl)
	ctx.Next()
}
