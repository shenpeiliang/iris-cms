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
	Click       uint
	Paixu       uint
	IsShow      byte `gorm:"column:is_show"`
}

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

//分页数据
func (a Article) Page(where map[string]interface{}, offset, limit uint) ([]Article, error) {
	var (
		data []Article
	)

	util.DB.Where(where).Offset(offset).Limit(limit).Find(&data)

	if len(data) < 1 {
		return data, errors.New("记录不存在")
	}

	return data, nil
}

//统计数
func (a Article) Count(where map[string]interface{}) (uint, error) {

	var (
		count uint
	)
	util.DB.Model(Article{}).Where(where).Count(&count)

	return count, nil
}
