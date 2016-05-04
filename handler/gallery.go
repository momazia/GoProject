package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

// Login handler
func GalleryHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/gallery.html"))
	// Fetching user's data
	pt := ProfileTemp{
		User:   util.GetUser(req),
	}
	err := tpl.Execute(res, GetAPlusTemplateHeader(req, pt))
	log.LogError(err)
}
