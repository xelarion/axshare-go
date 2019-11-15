package utils

import "golang.org/x/crypto/bcrypt"

func PasswordToBcrypt(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
