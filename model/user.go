package model

import (
	"cms/util"
	"errors"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"column:user_name"`
	Password string
	Dateline uint
	IsShow   byte
	AddTime  uint
}

//保存数据
func (u User) Save(data User) (ret bool, err error) {

	if data.ID > 0 {
		//更新
		err = util.DB.Model(User{}).Where(&User{
			ID: data.ID,
		}).Update(data).Error
	} else {
		//新增
		err = util.DB.Model(User{}).Save(&data).Error
	}

	if err != nil {
		return false, err
	}

	ret = true

	return
}

//指定字段更新
func (u User) Update(data, where map[string]interface{}) (ret bool, err error) {

	err = util.DB.Model(User{}).Where(where).Update(data).Error

	if err != nil {
		return false, err
	}

	ret = true

	return
}

//获取一条记录
func (u User) Get(id uint) (User, error) {
	var data User
	util.DB.Where(&User{
		ID: id,
	}).First(&data)

	if data.ID < 1 {
		return data, errors.New("记录不存在")
	}

	return data, nil
}

//获取指定一个用户信息
func (u User) GetUser(where User) User {
	var data User
	util.DB.Where(where).First(&data)

	return data
}

//分页数据
func (u User) Page(where map[string]interface{}, offset, limit uint, order string) ([]User, error) {
	var (
		data []User
	)

	db := util.DB

	if len(where) > 0 {
		for key, value := range where {
			db = db.Where(key, value)
		}
	}

	db.Offset(offset).Limit(limit).Order(order).Find(&data)

	if len(data) < 1 {
		return data, errors.New("记录不存在")
	}

	return data, nil
}

//统计数
func (u User) Count(where map[string]interface{}, order string) (uint, error) {

	var (
		count uint
	)

	db := util.DB
	if len(where) > 0 {
		for key, value := range where {
			db = db.Where(key, value)
		}
	}

	db.Model(User{}).Order(order).Count(&count)

	return count, nil
}

//删除记录
func (u User) Delete(id uint) (ret bool, err error) {
	err = util.DB.Model(User{}).Unscoped().Delete(&User{
		ID: id,
	}).Error

	if err != nil {
		return
	}

	return true, nil
}
