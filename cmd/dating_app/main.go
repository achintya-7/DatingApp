package main

import (
	"github.com/achintya-7/dating-app/config"
	"github.com/achintya-7/dating-app/utils"
)

func init() {
	config.LoadConfig()
}

func main() {
	utils.ApplyMigrations()
}