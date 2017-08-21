package util

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
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

//StartFileLog initialises the logging function
func StartFileLog(localLog *log.Logger) (*log.Logger, *os.File) {
	if _, err := os.Stat("./log/"); os.IsNotExist(err) {
		os.Mkdir("./log/", 0666)
	}
	f, err := os.OpenFile("./log/Main.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		localLog.Fatalf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)

	localLog.SetOutput(mw)
	localLog.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
	return localLog, f
}
