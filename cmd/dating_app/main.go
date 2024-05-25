package main

import (
	"github.com/achintya-7/dating-app/config"
	"github.com/achintya-7/dating-app/logger"
	"github.com/achintya-7/dating-app/utils"
)

func init() {
	logger.LoadLogger()
	config.LoadConfig()
	utils.ApplyMigrations()
}

func main() {
	
}
	
