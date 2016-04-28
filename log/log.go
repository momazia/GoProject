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
