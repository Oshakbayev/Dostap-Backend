package main

import (
	_ "github.com/lib/pq"
	configPkg "hellowWorldDeploy/cmd/config"
	"hellowWorldDeploy/logs"
	"hellowWorldDeploy/pkg"
	"log"
	"os"
)

func main() {

	config := configPkg.CreateConfig()
	if err := configPkg.ReadConfig("cmd/config/devConfig.json", config); err != nil {
		log.Fatal(err)
	}
	logFile := logs.CreateLogFile()
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(logFile)
	logger := logs.NewLogger(logFile)
	server := pkg.InitServer(config, logger)
	err := server.StartServer()
	if err != nil {
		logger.Printf("Emergency:Server crushed!!!")
		log.Fatal(err)
	}

}
