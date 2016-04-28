package main

import (
	"html/template"
	"log"
	"net/http"
)

const BUCKET_NAME = "gotraining-1271.appspot.com"

func init() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
}

// Signup handler
func signupHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/signup.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// Login handler
func loginHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/login.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// The main page handler
func indexHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/index.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
