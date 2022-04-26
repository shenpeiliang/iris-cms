package util

import (
	"flag"

	"github.com/kataras/iris/v12"
)

type Config struct {
}

//获取全部配置项
func (c Config) GetAll() iris.Configuration {
	var p string
	flag.StringVar(&p, "config", "./config/config.yml", "配置文件")
	flag.Parse()

	return iris.YAML(p)
}

//获取其他配置项
func (c Config) GetOther() map[string]interface{} {
	return c.GetAll().GetOther()
}
