package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/storage"
	"html/template"
	"net/http"
	"strings"
)

type GalleryTemp struct {
	Files []File
}

type File struct {
	Name      string
	Extension string
}

// Gallery handler
func GalleryHandler(res http.ResponseWriter, req *http.Request) {

	username := session.GetUser(req).Email

	if req.Method == "POST" {
		file, header, err := req.FormFile("file")
		log.LogErrorWithMsg("Cannot read the file from the request", err)
		if err == nil {
			err = storage.Store(req, username, header.Filename, file)
			log.LogErrorWithMsg("Cannot store the uploaded file", err)
		}
	}
	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/gallery.html"))

	// Getting user's list of file
	fileList, err := storage.RetrieveFileList(req, username)
	log.LogErrorWithMsg("Cannot get user's list of files", err)

	files := createFiles(fileList)

	gt := GalleryTemp{
		Files: files,
	}

	err = tpl.Execute(res, GetAPlusTemplateHeader(req, gt))
	log.LogError(err)
}

// Converts list of files into File
func createFiles(fileList []string) []File {
	var files []File
	for _, file := range fileList {
		str := strings.Split(file, ".")
		files = append(files, File{Name: strings.Join(str[:len(str)-1], "."), Extension: str[len(str)-1]})
	}
	return files
}
