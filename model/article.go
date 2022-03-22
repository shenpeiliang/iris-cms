package model

import (
	"cms/util"
	"errors"
)

type Article struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
	Content     string
	Img         string
	AddTime     uint `gorm:"column:add_time"`
	Dateline    uint
}

var ArticleModel = new(Article)

//获取一条记录
func (a Article) Get(id uint) (Article, error) {
	var data Article
	util.DB.Where(&Article{
		ID: id,
	}).First(&data)

	if data.ID < 1 {
		return data, errors.New("记录不存在")
	}

	return data, nil
}

//获取
func (a Article) Page(offset, limit uint) ([]Article, uint, error) {
	var (
		data  []Article
		count uint
	)
	util.DB.Offset(offset).Limit(limit).First(&data)

	util.DB.Count(&count)

	if len(data) < 1 {
		return data, count, errors.New("记录不存在")
	}

	return data, count, nil
}
