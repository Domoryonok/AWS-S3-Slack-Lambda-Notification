package utils

import "log"

func HandleError(err error, reason string) {
	if err != nil {
		log.Fatalf("%s error: %v", reason, err)
	}
}
