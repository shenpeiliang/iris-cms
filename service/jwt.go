package service

import (
	"github.com/iris-contrib/middleware/jwt"
)

//密钥
var mySecret = []byte("My Secret")

//初始化
func InitJWT() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		Expiration:    true,                   //超时检查
		SigningMethod: jwt.SigningMethodES256, //签名方法
	})
}

//获取token值
func GetJWTToken(data map[string]interface{}) (tokenString string, err error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))

	tokenString, err = token.SignedString(mySecret)

	return

}

//拆解token获取原数据
func GetJWTData(tokenString interface{}) map[string]interface{} {
	return tokenString.(*jwt.Token).Claims.(jwt.MapClaims)
}
