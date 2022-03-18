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
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		function.GetEnv("USER", "root"),
		function.GetEnv("PASSWORD", "root"),
		function.GetEnv("DATABASE", ""),
		function.GetEnv("CHARSET", "utf8"),
	))
	defer DB.Close()
}
