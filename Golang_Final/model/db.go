package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "DaU178007"
	dbname   = "Project"
)

func connect() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sql.Open("postgres", psqlconn)
}

func Query() []Product {
	db, err := connect()
	CheckError(err)
	defer db.Close()

	queryData := `select * from "products" order by id`
	a, e := db.Query(queryData)
	CheckError(e)
	defer a.Close()
	res := []Product{}
	for a.Next() {
		var p Product
		err = a.Scan(&p.Id, &p.Name, &p.Model, &p.Price, &p.Characteristic, &p.Size, &p.TotalRating, &p.PhotoUrl)
		CheckError(err)
		res = append(res, p)
	}
	return res

}

func find(u User) User {
	db, err := connect()
	CheckError(err)
	defer db.Close()
	findData := `select name, surname, login from "users" where login = $1 and password = $2`
	a, e := db.Query(findData, u.Login, u.Password)
	CheckError(e)
	defer a.Close()

	res := User{}
	for a.Next() {
		var u User
		err = a.Scan(&u.Name, &u.Surname, &u.Login)
		CheckError(err)
		res = u
	}
	return res

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
