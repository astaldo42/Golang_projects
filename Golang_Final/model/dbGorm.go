package model

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//)
//
//var (
//	db *gorm.DB
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "golang"
//)
//
//func Connect() {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	d, err := gorm.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	db = d
//}
//
//func GetDB() *gorm.DB {
//	return db
//}
//
////	func query() []Product {
////		var Products []Product
////		db.Order("id").Find(&Products)
////		return Products
////	}
//func signUp(user User) {
//	if len(user.name) < 20 || len(user.surname) < 20 && isAlpha(user.name) && isAlpha(user.surname) && checker(user.password) {
//		user.password = hashPassword(user)
//		//db.AutoMigrate(&User{})
//
//		newUser := &User{
//			name:     user.name,
//			surname:  user.surname,
//			login:    user.login,
//			password: user.password,
//		}
//
//		db.Create(newUser)
//	}
//}
