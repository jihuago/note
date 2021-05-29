package logging

import (
	"log"
	"os"
)

type Level int

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func init()  {
	filePath := getLogFilePath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}






















