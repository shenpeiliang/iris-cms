package util

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func init() {

	//初始化配置
	config := initMysqlConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config["User"],
		config["Password"],
		config["Host"],
		config["Port"],
		config["Database"],
		config["Charset"],
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config["TablePrefixt"].(string), //表前缀
			SingularTable: true,                            //禁用复数表名
		},
	})
	if err != nil {
		panic("数据库连接错误：" + err.Error())
	}

	db, err := DB.DB()
	if err != nil {
		panic("数据库连接错误：" + err.Error())
	}

	//连接池配置
	if _, has := config["MaxOpenConns"]; has {

		//连接池最大连接数
		db.SetMaxOpenConns(config["MaxOpenConns"].(int))

		//连接池最大允许的空闲连接数
		db.SetMaxIdleConns(config["MaxIdleConns"].(int))

		//SetConnMaxLifetime 设置了连接可复用的最大时间。
		db.SetConnMaxLifetime(time.Hour)

	} else {

		//数据库关闭（使用连接池禁止关闭）
		defer db.Close()

	}

}

//初始化配置项
func initMysqlConfig() map[string]interface{} {
	//配置
	config := Config{}.GetOther()

	//默认值
	rConfig := map[string]interface{}{
		"Host":         "127.0.0.1",
		"Port":         "3306",
		"User":         "root",
		"Password":     "root",
		"Charset":      "UTF8",
		"TablePrefixt": "",
	}

	//是否有配置项
	if c, has := config["Mysql"]; has {
		item := c.(map[string]interface{})
		if v, has := item["Host"]; has {
			rConfig["Host"] = v.(string)
		}

		if v, has := item["Port"]; has {
			rConfig["Port"] = v.(string)
		}

		if v, has := item["User"]; has {
			rConfig["User"] = v.(string)
		}

		if v, has := item["Password"]; has {
			rConfig["Password"] = v.(string)
		}

		if v, has := item["Database"]; has {
			rConfig["Database"] = v.(string)
		}

		if v, has := item["Charset"]; has {
			rConfig["Charset"] = v.(string)
		}

		if v, has := item["TablePrefixt"]; has {
			rConfig["TablePrefixt"] = v.(string)
		}

		if v, has := item["MaxOpenConns"]; has {
			rConfig["MaxOpenConns"] = v.(int)
		}

		if v, has := item["MaxIdleConns"]; has {
			rConfig["MaxIdleConns"] = v.(int)
		}

	}

	return rConfig
}
