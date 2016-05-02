package main

import (
	"github.com/momazia/GoProject/handler"
	"github.com/momazia/GoProject/session"
	"net/http"
)

func init() {
	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	// Setting the handlers
	http.Handle("/", handlerFilter(http.HandlerFunc(handler.IndexHandler)))
	http.Handle("/login", handlerFilter(http.HandlerFunc(handler.LoginHandler)))
	http.Handle("/signup", handlerFilter(http.HandlerFunc(handler.SignupHandler)))
}

// The filter is called whenever a handler is marked to go through this filter.
// The filter takes care of session management.
func handlerFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// Handling the session
		if !session.Handle(responseWriter, request) {
			http.Redirect(responseWriter, request, request.URL.Path, http.StatusMovedPermanently)
			return
		}
		handler.ServeHTTP(responseWriter, request)
	})
}
