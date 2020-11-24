package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plainText string) string {
	cipherText, err := bcrypt.GenerateFromPassword([]byte(plainText), 5)
	if err != nil {
		panic(err)
	}
	return string(cipherText)
}

func VerifyPassword(cipherText, plainText string) bool {
	return bcrypt.CompareHashAndPassword([]byte(cipherText), []byte(plainText)) == nil
}
