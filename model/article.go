package model

import "cms/util"

type Article struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
	Content     string
	Img         string
	AddTime     uint `gorm:"column:add_time"`
	Dateline    uint
}

//获取一条记录
func (a Article) Get(id uint) Article {
	var data Article
	util.DB.Where(&Article{
		ID: id,
	}).First(&data)

	return data
}
