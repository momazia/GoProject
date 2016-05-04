package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/storage"
	"io"
	"net/http"
)

// Download handler
func DownloadHandler(res http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Query().Get("fileName")
	rc, err := storage.Retrieve(req, session.GetUser(req).Email, fileName)
	log.LogErrorWithMsg("Cannot retrieve the file name given", err)
	if err == nil {
		io.Copy(res, rc)
		defer rc.Close()
	}
}
