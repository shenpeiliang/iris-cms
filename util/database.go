package util

import (
	"cms/function"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		function.GetEnv("USER", "root"),
		function.GetEnv("PASSWORD", "root"),
		function.GetEnv("HOST", "localhost"),
		function.GetEnv("PORT", "3306"),
		function.GetEnv("DATABASE", "xqw100"),
		function.GetEnv("CHARSET", "utf8"),
	)

	//修改默认表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return function.GetEnv("TABLE_PREFIX", "") + defaultTableName
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
	db := DB.DB()
	//连接池最大连接数
	db.SetMaxOpenConns(100)
	//连接池最大允许的空闲连接数
	db.SetMaxIdleConns(20)

	//数据库关闭（使用连接池需要禁止关闭）
	//defer DB.Close()
}
