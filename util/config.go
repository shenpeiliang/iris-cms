package util

import "github.com/kataras/iris/v12"

type Config struct {
}

//获取全部配置项
func (c Config) GetAll() iris.Configuration {
	return iris.YAML("./config/config.yml")
}

//获取其他配置项
func (c Config) GetOther() map[string]interface{} {
	return c.GetAll().GetOther()
}
