package logs

import (
	"log"
	"os"
)

func CreateLogFile() *os.File {
	logFile, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return logFile
}

func NewLogger(logFile *os.File) *log.Logger {
	return log.New(logFile, "---New log line---", log.Ldate|log.Ltime)
}
