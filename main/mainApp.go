package main

import (
	"github.com/momazia/GoProject/handler"
	"net/http"
)

const BUCKET_NAME = "gotraining-1271.appspot.com"

func init() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
}
