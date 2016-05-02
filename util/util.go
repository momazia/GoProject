package util

import (
	"github.com/momazia/GoProject/datastore"
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/memcache"
	"github.com/momazia/GoProject/session"
	"net/http"
)

// Stores the user in memcache and data store
func SaveUser(req *http.Request, u datastore.User) {
	// Storing into memcache
	err := memcache.Store(u.Email, session.User{Email: u.Email}, req)
	log.LogErrorWithMsg("Cannot store the user into memcache:", err)
	// Storing into datastore
	err = datastore.Store(req, datastore.KIND_USER, u)
	log.LogErrorWithMsg("Cannot store the user into datastore:", err)
}
