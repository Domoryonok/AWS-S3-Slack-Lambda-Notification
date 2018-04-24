package utils

import (
	"log"
	"runtime"
	"strings"
)

func getReadebleLogPath(path string) string {
	var splitPath []string =  strings.Split(path, "/")
	return strings.Join(splitPath[len(splitPath) - 3:], "/")
}

func HandleError(err error, reason string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = getReadebleLogPath(file)

		log.Fatalf("[.../%s:%d] - %s error : %v", file, line, reason, err)
	}
}
