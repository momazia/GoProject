package util

import (
	"github.com/momazia/GoProject/datastore"
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/memcache"
	"github.com/momazia/GoProject/user"
	"net/http"
)

// Stores the user in memcache and data store
func SaveUser(req *http.Request, u user.User) {
	// Storing into memcache
	err := memcache.Store(u.Email, u, req)
	log.LogErrorWithMsg("Cannot store the user into memcache:", err)
	// Storing into datastore
	var model = datastore.Model{
		ID:   u.Email,
		Data: u,
	}
	datastore.Store(req, datastore.KIND_USER, model)
	log.LogErrorWithMsg("Cannot store the user into datastore:", err)
}
