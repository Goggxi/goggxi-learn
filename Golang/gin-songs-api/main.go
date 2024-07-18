package main

import (
	"fmt"
	"gin-songs-api/config"
	_ "gin-songs-api/docs"
	"gin-songs-api/routes"
	"gin-songs-api/utils"
	"github.com/gin-gonic/gin"
)

// @title Songs API
// @version 1
// @description This is a simple songs API

// @contact.name Moh Rifkan
// @contact.email goggxi@gmail.com
// @contact.uRL https://github.com/Goggxi

// @host localhost:3000
// @BasePath /api/v1
func main() {
	cfg := utils.LoadConfig()

	config.InitDB(cfg)
	defer config.DB.Close()

	r := gin.Default()
	routes.SetupRoutes(r, config.DB)
	err := r.Run(cfg.ServerAddress)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
