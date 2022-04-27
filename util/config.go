package util

import (
	"flag"

	"github.com/kataras/iris/v12"
)

//配置信息
var globalConfig *iris.Configuration

type Config struct {
}

//配置信息
func (c Config) New() iris.Configuration {
	if globalConfig != nil {
		return *globalConfig
	}

	return c.GetAll()
}

//获取全部配置项
func (c Config) GetAll() iris.Configuration {
	var p string
	//是否被解析过
	if flag.Parsed() {
		p = flag.Arg(0)
		//fmt.Printf("输入：%v", flag.Args())
		if p == "" {
			p = "./config/config.yml"
		}
	} else {
		flag.StringVar(&p, "config", "./config/config.yml", "配置文件")
		flag.Parse()
	}

	config := iris.YAML(p)

	globalConfig = &config

	return config

}

//获取其他配置项
func (c Config) GetOther() map[string]interface{} {
	return c.New().GetOther()
}
