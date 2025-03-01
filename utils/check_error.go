package utils

import (
	"log"
	"runtime"
	"time"
)

// CheckError provides better error logging with function name and timestamp
func CheckError(err error) {
	if err != nil {
		// Get caller function details
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()

		// Format the error message
		log.Fatalf("[%s] ❌ ERROR in %s: %v", time.Now().Format("2006-01-02 15:04:05"), funcName, err)
	}
}
