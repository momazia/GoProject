package session

import (
	"github.com/momazia/GoProject/log"
	"net/http"
)

// Holds the UUIDs who are logged in successfully. This session is not invalidated at any time.
var sessions []string

// Handles session related operation. If the client does not have a cookie session set, it will create a new UUID
// and sets that on the cookie and it will register it in the session. If the client does have a valid session
// (by looking up the UUID coming from cookie in the session), then it will not do anything. Otherwise, it will
// assign a new session to the user again.
func Handle(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("Session handler is called")
}
