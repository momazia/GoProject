package handler

import (
	"github.com/momazia/GoProject/log"
	"html/template"
	"github.com/momazia/GoProject/session"
	"net/http"
)

// Signup handler
func SignupHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/signup.html"))
	err := tpl.Execute(res, session.GetUser(req).Name)
	log.LogError(err)
}
