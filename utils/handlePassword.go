package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// 使用bcrypt库的GenerateFromPassword函数进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败")
		fmt.Println(err)
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, inputPassword string) error {
	// 使用bcrypt库的CompareHashAndPassword函数比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}
