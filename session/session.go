package session

import (
	"fmt"
	"github.com/momazia/GoProject/log"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
)

type session struct {

	// Holds the UUIDs who are logged in successfully. This session is not invalidated at any time.
	SessionIds []string
}

var s *session

func init() {
	s = &session{}
}

const SESSION_ID = "SESSION-ID"

// Handles session related operation. If the client does not have a cookie session set, it will create a new UUID
// and sets that on the cookie and it will register it in the session. If the client does have a valid session
// (by looking up the UUID coming from cookie in the session), then it will not do anything. Otherwise, it will
// assign a new session to the user again.
func Handle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)

	//	if isUserInSession(responseWriter, request) {
	//		log.Println("User is in session.")
	//	} else {
	//		log.Println("User is not in session, setting the SESSION-ID ...")
	//		createSession(responseWriter)
	//	}
}

// Creates a session by creating a new UUID and setting it on cookie and sessions.
func createSession(res http.ResponseWriter) {
	newUuid, err := uuid.NewV4()
	sessionId := newUuid.String()
	log.LogError(err)
	s.SessionIds = append(s.SessionIds, sessionId)
	createCookie(&res, SESSION_ID, sessionId)
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
func isUserInSession(res http.ResponseWriter, req *http.Request) bool {
	log.Println("Session")
	log.Println(s.SessionIds)
	sessionIdCookie, err := req.Cookie(SESSION_ID)
	if err != nil {
		log.Println("Error reading SESSIONID:" + err.Error())
		return false
	}
	for _, session := range s.SessionIds {
		if session == sessionIdCookie.Value {
			return true
		}
	}
	return false
}
