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

func StartFileLog(localLog *log.Logger) {
	if _, err := os.Stat("./log/"); os.IsNotExist(err) {
		os.Mkdir("./log/", 0666)
	}
	f, err := os.OpenFile("./log/Main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	localLog.SetOutput(f)
}
