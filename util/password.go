package util

import "golang.org/x/crypto/bcrypt"

//密码校验
//hashedPassword加密的密码
//表单密码
func CompareHashAndPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

//加密密码
func EncryptedPassword(password string) (hashedPassword string, err error) {
	var pwd []byte
	if pwd, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return
	}
	return string(pwd), err
}
