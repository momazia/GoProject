package storage

import (
	"github.com/momazia/GoProject/log"
	"google.golang.org/appengine"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
)

const BUCKET_NAME = "gotraining-1271.appspot.com"

// Saves the given file for the given name and content under the username.
func Store(req *http.Request, userName, fileName string, file io.Reader) error {

	fileName = userName + "/" + fileName

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	defer client.Close()

	writer := client.Bucket(BUCKET_NAME).Object(fileName).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{
		storage.AllUsers,
		storage.RoleReader}}

	io.Copy(writer, file)
	writer.Close()
	return err
}

// Retrieves the file stored for the given name
func Retrieve(req *http.Request, fileName string) (io.ReadCloser, error) {
	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	log.LogErrorWithMsg("Cannot create a new client", err)
	defer client.Close()
	return client.Bucket(BUCKET_NAME).Object(fileName).NewReader(ctx)
}
