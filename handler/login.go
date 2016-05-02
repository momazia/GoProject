package handler

import (
	"github.com/momazia/GoProject/log"
	"html/template"
	"net/http"
)

// Login handler
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		// Validation comes here
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/login.html"))
	err := tpl.Execute(res, nil)
	log.LogError(err)

}
