package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/storage"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

// Login handler
func GalleryHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		file, header, err := req.FormFile("file")
		log.LogErrorWithMsg("Cannot read the file from the request", err)
		err = storage.Store(req, session.GetUser(req).Email, header.Filename, file)
		log.LogErrorWithMsg("Cannot store the uploaded file", err)
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/gallery.html"))
	// Fetching user's data
	pt := ProfileTemp{
		User: util.GetUser(req),
	}

	err := tpl.Execute(res, GetAPlusTemplateHeader(req, pt))
	log.LogError(err)
}
