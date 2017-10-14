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
func StartFileLog(localLog *log.Logger) (logger *log.Logger, logFile *os.File, err error) {
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/log/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	f, err := os.OpenFile(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/Morpheus.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		localLog.Printf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)

	localLog.SetOutput(mw)
	localLog.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
	logger = localLog
	logFile = f
	return
}
