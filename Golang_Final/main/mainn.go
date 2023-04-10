package main

import (
	c "awesomeProject/console"
	"net/http"
)

func main() {

	http.HandleFunc("/", c.HttpGetProduct)
	http.HandleFunc("/login", c.HttpSignIn)
	http.HandleFunc("/signUp", c.HttpSignUp)
	http.HandleFunc("/search", c.HttpSearch)
	http.HandleFunc("/login/", c.AuthHandler)
	http.HandleFunc("/sortByPrice", c.HttpSortByPrice)
	http.HandleFunc("/sortByRating", c.HttpSortByRating)

	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}

}
