package handler

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/session"
	"github.com/momazia/GoProject/util"
	"html/template"
	"net/http"
)

// Login handler
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	invalidUser := false

	if req.Method == "POST" {
		email := req.FormValue("username")
		if isValidUser(email, req.FormValue("password"), req) {
			// Set the session
			session.CreateSession(&res, req, session.User{Email: email})
			// Redirecting the user to profile page.
			http.Redirect(res, req, URL_PROFILE, http.StatusFound)
			return
		} else {
			// Invalid User
			invalidUser = true
		}
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/login.html"))
	err := tpl.Execute(res, invalidUser)
	log.LogError(err)

}

// Validates to see if the username and password given correct or not
func isValidUser(username, password string, req *http.Request) bool {
	u := util.GetUserWithEmail(username, req)
	return u.Email == username && u.Password == Encrypt(password)
}
