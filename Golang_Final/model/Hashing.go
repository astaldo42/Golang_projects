package model

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func isAlpha(s string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if isAlpha(s) {
		return true
	}
	return false
}

func checker(password string) bool {
	for _, char := range password {
		if char == ' ' {
			return false
		}
	}
	if !isAlpha(password) {
		return false
	}
	return true
}

func hashPassword(user User) string {
	h := md5.Sum([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h)
	return user.Password
}
