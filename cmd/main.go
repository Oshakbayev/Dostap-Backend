package main

import (
	configPkg "hellowWorldDeploy/cmd/config"
	"hellowWorldDeploy/logs"
	"hellowWorldDeploy/pkg"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if _, err := os.ReadFile(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	

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
