package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"html/template"
	"github.com/momazia/GoProject/session"
	"net/http"
)

// Login handler
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/login.html"))
	err := tpl.Execute(res, session.GetUser(req).Name)
	log.LogError(err)
}
