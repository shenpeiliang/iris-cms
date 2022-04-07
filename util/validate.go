package util

import (
	"reflect"
	"strconv"

	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

//自定义验证规则
func StrLenFunc(fl validator.FieldLevel) bool {
	param, _ := strconv.Atoi(fl.Param())

	valueLen := len(fl.Field().String())

	if valueLen < param {
		return false
	}
	return true
}

//处理表单验证错误
func ValidateErrHandle(ctx iris.Context, obj interface{}, err error) {
	//是否空值
	if _, ok := err.(*validator.InvalidValidationError); ok {
		Response.Fail(ctx, err.Error())

		return
	}

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		fieldName := e.Field()
		//反射获取其他标签信息
		field, ok := reflect.TypeOf(obj).FieldByName(fieldName)
		errInfo := field.Tag.Get("error-" + e.Tag())

		if ok {
			Response.Fail(ctx, errInfo)
			return
		}
	}

	Response.Fail(ctx, err.Error())
	return
}
