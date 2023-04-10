package model

type User struct {
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Login    string `json:"login" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type Client struct {
	User
	//liked []Product
}

func (u User) insert() {
	db, err := connect()
	CheckError(err)
	defer db.Close()
	insertData := `insert into "users"("name", "surname","login","password") values($1,$2,$3,$4)`
	_, e := db.Exec(insertData, u.Name, u.Surname, u.Login, u.Password)
	CheckError(e)

}

func (u User) getInform() (name string, surname string) {
	name = u.Name
	surname = u.Surname
	return
}
