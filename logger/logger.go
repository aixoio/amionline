package logger

import (
	"log"
	"os"
)

var error log.Logger
var warn log.Logger
var info log.Logger

func Init() {
	error = *log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	warn = *log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	info = *log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
}

func Error() *log.Logger {
	return &error
}

func Warn() *log.Logger {
	return &warn
}

func Info() *log.Logger {
	return &info
}
