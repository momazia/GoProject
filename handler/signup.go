package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/user"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

// Signup handler
func SignupHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		// Validatation comes here

		// Create the user
		var user = user.User{
			Email:     req.FormValue("email"),
			Password:  req.FormValue("password1"),
			FirstName: req.FormValue("fName"),
			LastName:  req.FormValue("lName"),
		}

		// Store the user
		util.SaveUser(req, user)

		// Create session
		session.CreateSession(&res, req, user)

		// Redirect the user to front page
		http.Redirect(res, req, URL_ROOT, http.StatusFound)
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/signup.html"))
	err := tpl.Execute(res, nil)
	log.LogError(err)
}
