package function

import (
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12/view"
)

//获取环境配置
func GetEnv(key string, def string) string {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	v := os.Getenv(key)

	if v == "" {
		return def
	}

	return v
}

//注册视图函数
func RegisterTemplateFun(tmpl *view.HTMLEngine) {
	//css文件
	tmpl.AddFunc("css", func(files ...string) template.HTML {
		var html string

		if len(files) < 1 {
			return template.HTML(html)
		}

		for _, file := range files {

			html += "<link  type='text/css' rel='stylesheet' href='" + file + "?v=" + GetEnv("STATIC_VERSION", "0.0.0") + "'/>"
		}

		return template.HTML(html)
	})

	//css文件
	tmpl.AddFunc("js", func(files ...string) template.HTML {
		var html string
		if len(files) < 1 {
			return template.HTML(html)
		}

		for _, file := range files {

			html += "<script type='text/javascript' src='" + file + "?v=" + GetEnv("STATIC_VERSION", "0.0.0") + "'></script>"
		}

		return template.HTML(html)
	})
}
