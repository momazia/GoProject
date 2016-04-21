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
}
// Login handler
func loginHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("login.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// The main page handler
func indexHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("index.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
