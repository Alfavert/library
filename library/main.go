package main

import (
	//"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	//_ "github.com/go-sql-driver/mysql"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "home_page", nil)
}
func signup_page(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/signup_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		logrus.Warnf("Got error in [handler], err: %v", err)
		_, _ = io.WriteString(w, fmt.Sprintf("Error: %v", err))
		return
	}

	tmpl.ExecuteTemplate(w, "signup_page", nil)
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/signup_page/", signup_page)
	//http.HandleFunc("/contacts/", contact_page)
	http.ListenAndServe(":3333", nil)
}

func main() {
	handleRequest()
}
