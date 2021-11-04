package utils

import (
	"fmt"
	"time"
)

// Init Time Format
const TIME_FORMAT string = "2006-01-02 15:04:05.000"

// Get Time
func getTime() string {
	return time.Now().Format(TIME_FORMAT)
}

// Set Logger
func SetLogger(logInfo ...interface{}) {
	fmt.Printf("%s  ->  %s\n", getTime(), logInfo)
}
