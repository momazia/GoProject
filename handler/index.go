package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"html/template"
	"net/http"
)

// The main page handler
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/index.html"))
	err := tpl.Execute(res, session.GetUser(req).Name)
	log.LogError(err)
}
