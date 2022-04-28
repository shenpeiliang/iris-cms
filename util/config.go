package util

import (
	"flag"
	"sync"

	"github.com/kataras/iris/v12"
)

//配置信息
var (
	globalConfig *iris.Configuration
	once         sync.Once
)

type Config struct {
}

//配置信息
func (c Config) New() iris.Configuration {
	once.Do(func() {
		c.loadConfig()
	})

	return *globalConfig
}

//加载配置
func (c Config) loadConfig() {
	var (
		configPath  string
		defaultPath = "./config/config.yml"
	)
	//是否被解析过
	if flag.Parsed() {
		configPath = flag.Arg(0)
		//fmt.Printf("输入：%v", flag.Args())
		if configPath == "" {
			configPath = defaultPath
		}
	} else {
		flag.StringVar(&configPath, "config", defaultPath, "配置文件")
		flag.Parse()
	}

	config := iris.YAML(configPath)

	globalConfig = &config

}

//获取其他配置项
func (c Config) GetOther() map[string]interface{} {
	return c.New().GetOther()
}
