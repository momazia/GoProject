package session

import (
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/memcache"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

const SESSION_ID = "SESSION-ID"

// Handles session related operation. If the client does not have a cookie session set, it will create a new UUID
// and sets that on the cookie and it will register it in the session. If the client does have a valid session
// (by looking up the UUID coming from cookie in the session), then it will not do anything. Otherwise, it will
// assign a new session to the user again.
func Handle(res http.ResponseWriter, req *http.Request) {

	if isUserInSession(req) {
		log.Println("User is in session.")
	} else {
		log.Println("User is not in session, setting the SESSION-ID ...")
		createSession(&res, req)
	}
}

// Creates a session by creating a new UUID and setting it on cookie and sessions.
func createSession(res *http.ResponseWriter, req *http.Request) {
	newUuid, err := uuid.NewV4()
	sessionId := newUuid.String()
	log.LogError(err)
	memcache.Store(sessionId, true, req) // Hard coding true value
	createCookie(res, SESSION_ID, sessionId)
}

// Creates a cookie for the given name and value and sets it on the response
func createCookie(res *http.ResponseWriter, cookieName, cookieValue string) {
	// Setting the cookie
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		HttpOnly: true,
	}
	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

// Checks to see if the user is logged in by looking at the sessionID stored on the request cookie
func isUserInSession(req *http.Request) bool {
//	sessionIdCookie, err := req.Cookie(SESSION_ID)
//	if err != nil {
//		log.Println("Error reading SESSIONID:" + err.Error())
//		return false
//	}
	var isLoggedIn string
	// Retrieve the item from memcache
	memcache.Retrieve("03ec8f2c-a224-4710-511e-ecc5fb946918", req, &isLoggedIn)
	log.Println("session: " + "03ec8f2c-a224-4710-511e-ecc5fb946918")
	log.Println("isLoggedIn: " + isLoggedIn)
	if isLoggedIn != "" {
		return true
	}
	return false
}
