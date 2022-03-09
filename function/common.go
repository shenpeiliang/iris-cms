package function

import (
	"os"

	"github.com/kataras/iris/v12/view"
)

//获取环境配置
func GetEnv(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

//注册视图函数
func RegisterTemplateFun(tmpl *view.HTMLEngine) {
	//css文件
	tmpl.AddFunc("css", func(files ...string) (html string) {
		if len(files) < 1 {
			return
		}

		for _, file := range files {

			html += "<link  type='text/css' rel='stylesheet' href='" + file + "?v=" + GetEnv("STATIC_VERSION", "0.0.0") + "'/>\n"
		}

		return html
	})
}
