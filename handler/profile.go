package handler

import (
	"github.com/momazia/GoProject/datastore"
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

type ProfileTemp struct {
	User   datastore.User
	Errors []string
}

// Profile handler
func ProfileHandler(res http.ResponseWriter, req *http.Request) {

	var errs []string

	if req.Method == "POST" {
		// Validation comes here
		pass1 := req.FormValue("password1")
		pass2 := req.FormValue("password2")
		errs = validateProfile(pass1, pass2)
		if errs == nil {
			// Saving the form
			du := datastore.User{
				Email:     session.GetUser(req).Email,
				Password:  pass1,
				FirstName: req.FormValue("fname"),
				LastName:  req.FormValue("lname"),
			}
			// Store the user
			util.SaveUser(req, du)

			// Redirect the user to front page
			http.Redirect(res, req, URL_ROOT, http.StatusFound)
		}
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/profile.html"))
	// Fetching user's data
	pt := ProfileTemp{
		User:   util.GetUser(req),
		Errors: errs,
	}
	err := tpl.Execute(res, GetAPlusTemplateHeader(req, pt))
	log.LogError(err)
}

// Validates to see if the form data is valid
func validateProfile(pass1, pass2 string) []string {
	var errs []string
	if pass1 == "" || pass2 == "" {
		errs = append(errs, "Password is mandatory")
	}
	if pass1 != pass2 {
		errs = append(errs, "Passwords do not match")
	}
	return errs
}
