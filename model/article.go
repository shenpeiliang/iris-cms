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

//保存数据
func (a Article) Save(data Article) (ret bool, err error) {

	if data.ID > 0 {
		//更新
		err = util.DB.Model(Article{}).Where(&Article{
			ID: data.ID,
		}).Update(data).Error
	} else {
		//新增
		err = util.DB.Model(Article{}).Save(&data).Error
	}

	if err != nil {
		return false, err
	}

	ret = true

	return
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

	db := util.DB

	if len(where) > 0 {
		for key, value := range where {
			db = db.Where(key, value)
		}
	}

	db.Offset(offset).Limit(limit).Find(&data)

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

	db := util.DB
	if len(where) > 0 {
		for key, value := range where {
			db = db.Where(key, value)
		}
	}

	db.Model(Article{}).Count(&count)

	return count, nil
}

//删除记录
func (a Article) Delete(id uint) (ret bool, err error) {
	err = util.DB.Unscoped().Delete(&Article{
		ID: id,
	}).Error

	if err != nil {
		return
	}

	return true, nil
}
