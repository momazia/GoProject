package main

import (
	hdl "github.com/momazia/GoProject/handler"
	"github.com/momazia/GoProject/session"
	"net/http"
)

func init() {
	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle(hdl.URL_STATIC+"/", http.StripPrefix(hdl.URL_STATIC, http.FileServer(http.Dir("."+hdl.URL_STATIC))))
	// Setting the handlers
	http.Handle(hdl.URL_ROOT, http.HandlerFunc(hdl.IndexHandler))
	http.Handle(hdl.URL_LOGIN, http.HandlerFunc(hdl.LoginHandler))
	http.Handle(hdl.URL_SIGN_UP, http.HandlerFunc(hdl.SignupHandler))
	http.Handle(hdl.URL_PROFILE, handlerFilter(http.HandlerFunc(hdl.ProfileHandler)))
	http.Handle(hdl.URL_LOGOUT, handlerFilter(http.HandlerFunc(hdl.LogoutHandler)))
}

// The filter is called whenever a handler is marked to go through this filter.
// The filter takes care of session management.
func handlerFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Handling the session
		if !session.Handle(responseWriter, request) {
			http.Redirect(responseWriter, request, hdl.URL_LOGOUT, http.StatusFound)
			return
		}
		handler.ServeHTTP(responseWriter, request)
	})
}
