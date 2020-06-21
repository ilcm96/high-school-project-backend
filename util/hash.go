package util

import "golang.org/x/crypto/bcrypt"

func HashPW(pw string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashedPW)
}
