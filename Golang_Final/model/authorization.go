package model

import (
	"fmt"
)

var Products = Query()
var Clientt Client

type Inserter interface {
	insert()
}

func SignUp(user User) {
	if len(user.Name) < 20 || len(user.Surname) < 20 && isAlpha(user.Name) && isAlpha(user.Surname) && checker(user.Password) {
		user.Password = hashPassword(user)
		Inserter.insert(user)
		fmt.Printf("Succes authorization \n")
	}
}
func SignIn(login string, password string) User {
	user := User{Login: login, Password: password}
	user.Password = hashPassword(user)
	user = find(user)

	return user
}
