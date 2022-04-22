package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	//修改默认表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config["TablePrefixt"].(string) + defaultTableName
	}

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("数据库连接错误：" + err.Error())
	}

	//调试模式
	DB.LogMode(true)

	//禁用表名复数形式
	DB.SingularTable(true)

	//连接池配置
	if _, has := config["MaxOpenConns"]; has {
		db := DB.DB()

		//连接池最大连接数
		db.SetMaxOpenConns(config["MaxOpenConns"].(int))

		//连接池最大允许的空闲连接数
		db.SetMaxIdleConns(config["MaxIdleConns"].(int))

	} else {

		//数据库关闭（使用连接池禁止关闭）
		defer DB.Close()

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
