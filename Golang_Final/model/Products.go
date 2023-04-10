package model

import "fmt"

type Product struct {
	Id             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Model          string `json:"model"`
	Price          int    `json:"price"`
	Characteristic string `json:"characteristic"`
	Size           string `json:"size"`
	comment        []Comment
	TotalRating    float32 `json:"totalRating"`
	PhotoUrl       string  `json:"photoUrl"`
}

type Comment struct {
	text   string
	rating float32
	owner  Client
}

func appendToStore(product Product) {
	Inserter.insert(product)
}

func (p Product) insert() {
	db, err := connect()
	CheckError(err)
	defer db.Close()
	insertData := `insert into "products"("id","name","model","price","characteristic","size","totalrating","photourl")values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, e := db.Exec(insertData, p.Id, p.Name, p.Model, p.Price, p.Characteristic, p.Size, p.TotalRating, p.PhotoUrl)
	CheckError(e)
}

func (p Product) addcomment(comment Comment) {
	fmt.Print(p.Id)
	db, err := connect()
	CheckError(err)
	defer db.Close()
	insertData := `insert into "comment" ("product_id","text","rating","commetator") values($1,$2,$3,$4)`
	_, e := db.Exec(insertData, p.Id, comment.text, comment.rating, comment.owner.Login)
	CheckError(e)
	p.setRating()
}

func (p *Product) getComment() {
	db, err := connect()
	CheckError(err)
	defer db.Close()
	query := `select text,rating,commetator from "comment" where product_id = $1`
	a, e := db.Query(query, p.Id)
	CheckError(e)
	defer a.Close()
	for a.Next() {
		var comments Comment
		err := a.Scan(&comments.text, &comments.rating, &comments.owner.Login)
		CheckError(err)
		p.comment = append(p.comment, comments)
	}
}

func printComment(p Product) {
	p.getComment()
	fmt.Printf("Comment for %s %s %s \n", p.Name, p.Model, p.Size)
	for _, val := range p.comment {
		fmt.Printf("%s: %s; rating:%v \n", val.owner.Login, val.text, val.rating)
	}
}

func (p Product) calcRating() float32 {
	p.getComment()
	var res float32 = 0.0
	for _, val := range p.comment {
		res += val.rating
	}
	return res / float32(len(p.comment))
}

func (p Product) setRating() {
	db, err := connect()
	CheckError(err)
	defer db.Close()
	setData := `update "products" set totalRating = $1 where id = $2`
	_, e := db.Exec(setData, p.calcRating(), p.Id)
	CheckError(e)
}
func getProducts(p []Product) {
	for a, val := range p {
		fmt.Printf("%v %s %s  %s %vtg rating: %v  \n", a, val.Name, val.Model, val.Size, val.Price, val.TotalRating)
	}
}
func SortByPrice(products []Product) []Product {
	qSort(products, 0, len(products)-1, 1)
	//getProducts(products)
	return products
}

func SortByRating(products []Product) []Product {
	qSort(products, 0, len(products)-1, 0)
	//getProducts(product)
	return products
}

func Search(q string) []Product {
	res := []Product{}
	for _, val := range Query() {
		if val.Name == q {
			res = append(res, val)
		}
	}
	return res
}
