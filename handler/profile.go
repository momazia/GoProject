package handler

import (
	"github.com/momazia/GoProject/log"
	"html/template"
	"net/http"
)

// Profile handler
func ProfileHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/profile.html"))
	err := tpl.Execute(res, GetAPlusTemplateHeader(req, nil))
	log.LogError(err)
}
