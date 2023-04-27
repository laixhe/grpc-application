package utils

import "golang.org/x/crypto/bcrypt"

// BcryptPasswordHash 进行加密
func BcryptPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// BcryptPasswordCheck 对比密码哈希值
func BcryptPasswordCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
