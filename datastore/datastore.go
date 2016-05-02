package datastore

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

const (
	KIND_USER string = "USER"
)

// Stores the given model for the the kind in data store
func Store(req *http.Request, kind string, model Model) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, kind, model.ID, 0, nil)
	_, err := datastore.Put(ctx, key, &model.Data)
	return err
}

// Retrieves the model passed using the ID inside of it for the kind given.
func Retrieve(req *http.Request, kind string, model *Model) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, kind, model.ID, 0, nil)
	return datastore.Get(ctx, key, &model)
}
