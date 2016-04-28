package handler

import (
	"github.com/momazia/GoProject/log"
	"html/template"
	"net/http"
)

// Login handler
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/login.html"))
	err := tpl.Execute(res, nil)
	log.LogError(err)
}
