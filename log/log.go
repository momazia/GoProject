package log

import (
	"log"
)

// Logs the error given into log
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Prints the given interface
func Println(v ...interface{}) {
	log.Println(v)
}
