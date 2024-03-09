package main

import (
	"fmt"
	_ "github.com/lib/pq"
	configPkg "hellowWorldDeploy/cmd/config"
	"hellowWorldDeploy/logs"
	"hellowWorldDeploy/pkg"
	"log"
	"os"
)

func main() {
	fmt.Println("Tpotentially final test of ci/cd 3.0 ")
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
		log.Fatal(err)
	}

}
