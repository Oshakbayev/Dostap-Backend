package logs

import (
	"log"
	"os"
)

// log dont work
func CreateLogFile() *os.File {
	logFile, err := os.Create("logs/logs.log")
	if err != nil {
		log.Fatal(err)
	}
	return logFile
}

func NewLogger(logFile *os.File) *log.Logger {
	return log.New(logFile, "---New log line---", log.Ldate|log.Ltime)
}
