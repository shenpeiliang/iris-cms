package service

import (
	"cms/util"
	"time"

	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

//数据库
func GetRedisDB() *redis.Database {
	r := redis.New(initRedisConfig())

	if r == nil {
		panic("缓存数据库连接错误")
	}

	return r
}

//初始化配置项
func initRedisConfig() redis.Config {
	//配置
	config := util.Config{}.GetOther()

	//默认配置
	rConfig := redis.DefaultConfig()

	//是否有配置项
	if c, has := config["Redis"]; has {
		item := c.(map[string]interface{})
		if v, has := item["Network"]; has {
			rConfig.Network = v.(string)
		}

		if v, has := item["Addr"]; has {
			rConfig.Addr = v.(string)
		}

		if v, has := item["Timeout"]; has {
			rConfig.Timeout = time.Duration(v.(int)) * time.Second
		}

		if v, has := item["MaxActive"]; has {
			rConfig.MaxActive = v.(int)
		}

		if v, has := item["Password"]; has {
			rConfig.Password = v.(string)
		}

		if v, has := item["Database"]; has {
			rConfig.Database = v.(string)
		}

		if v, has := item["Prefix"]; has {
			rConfig.Prefix = v.(string)
		}

	}

	return rConfig
}
