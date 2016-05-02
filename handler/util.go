package handler

import (
	"github.com/momazia/GoProject/session"
	"net/http"
)

type APlusTemplate struct {
	Header Header
	Data   interface{}
}

type Header struct {
	IsLoggedIn bool
}

// Creates a template for the given data by including the standard header to it.
func GetAPlusTemplateHeader(req *http.Request, data interface{}) APlusTemplate {
	return APlusTemplate{
		Header: Header{
			IsLoggedIn: session.GetUser(req).Email != "",
		},
		Data: data,
	}
}
