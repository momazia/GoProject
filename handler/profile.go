package handler

import (
	"github.com/momazia/GoProject/datastore"
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

// Profile handler
func ProfileHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		// Validation comes here

		// Saving the form
		du := datastore.User{
			Email:     session.GetUser(req).Email,
			Password:  req.FormValue("password1"),
			FirstName: req.FormValue("fname"),
			LastName:  req.FormValue("lname"),
		}
		// Store the user
		util.SaveUser(req, du)

		// Redirect the user to front page
		http.Redirect(res, req, URL_ROOT, http.StatusFound)
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/profile.html"))
	// Fetching user's data
	u := util.GetUser(req)
	err := tpl.Execute(res, GetAPlusTemplateHeader(req, u))
	log.LogError(err)
}
