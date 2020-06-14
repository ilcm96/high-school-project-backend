package db

import "golang.org/x/crypto/bcrypt"

func hashPW(pw string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashedPW)
}
