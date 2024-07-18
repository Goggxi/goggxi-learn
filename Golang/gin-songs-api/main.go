package main

import (
	"fmt"
	"gin-songs-api/config"
	"gin-songs-api/routes"
	"gin-songs-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := utils.LoadConfig()

	config.InitDB(cfg)
	defer config.DB.Close()

	r := gin.Default()
	routes.SetupRoutes(r, config.DB)
	err := r.Run(":3000")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
