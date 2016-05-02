package handler

import (
	"github.com/momazia/GoProject/log"
	"html/template"
	"net/http"
)

// The main page handler
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/index.html"))
	tmp := GetAPlusTemplateHeader(req)
	log.Println("isLoggedIn:")
	log.Println(tmp.Header.IsLoggedIn)
	err := tpl.Execute(res, tmp)
	log.LogError(err)
}
