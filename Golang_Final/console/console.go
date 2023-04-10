package console

import (
	m "awesomeProject/model"

	"html/template"
	"net/http"
)

func HttpSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		m.Clientt = m.Client{m.SignIn(email, password)}
		if m.Clientt.Name != "" {
			http.Redirect(w, r, "http://localhost:8181", http.StatusSeeOther)
		}
	}
	http.ServeFile(w, r, "static/authorization.html")
}

func HttpSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		client := m.Client{m.User{Name: name, Login: email, Password: password, Surname: surname}}
		m.SignUp(client.User)

		http.Redirect(w, r, "http://localhost:8181", http.StatusSeeOther)

	}
	http.ServeFile(w, r, "static/SignUp.html")
}

func HttpSearch(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/home.html")
	if r.Method == "POST" {
		m.Products = m.Search(r.FormValue("search"))
		tmpl.Execute(w, m.Products)
	}

}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the authorization page
	http.Redirect(w, r, "http://localhost:8181/login", http.StatusTemporaryRedirect)
}

func HttpGetProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/home.html")
	tmpl.Execute(w, m.Products)

}
func HttpSortByPrice(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/home.html")
	tmpl.Execute(w, m.SortByPrice(m.Products))
}
func HttpSortByRating(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/home.html")
	tmpl.Execute(w, m.SortByRating(m.Products))
}
