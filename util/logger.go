package util

import (
	"log"
	"os"
	"sync"
)

var logInstance *log.Logger
var once sync.Once

//Logger Exports a default Logger
func Logger() *log.Logger {
	once.Do(func() {
		logInstance = &log.Logger{}
	})
	return logInstance
}

func init() {
	f, err := os.OpenFile("log/Main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	localLog := Logger()

	localLog.SetOutput(f)
	localLog.Println("This is a test log entry")
}
