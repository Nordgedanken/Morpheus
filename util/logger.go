package util

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/shibukawa/configdir"
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
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, err := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/log/"); os.IsNotExist(err) {
		os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/", 0666)
	}
	f, err := os.OpenFile(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/Morpheus.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		localLog.Printf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)

	localLog.SetOutput(mw)
	localLog.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
	return localLog, f
}
