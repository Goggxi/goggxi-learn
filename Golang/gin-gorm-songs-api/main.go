package main

import (
	"gin-gorm-songs-api/configs"
	"gin-gorm-songs-api/models/entities"
	"gin-gorm-songs-api/utils"
	"gorm.io/gorm"
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
	db := configs.InitDB(cfg)

	err := db.Transaction(func(tx *gorm.DB) error {
		// Create
		artist := entities.Artist{Name: "John Doe"}
		tx.Create(&artist)

		// Read
		var readArtist entities.Artist
		tx.First(&readArtist, artist.ID)

		// Update
		tx.Model(&readArtist).Update("Name", "Jane Doe")

		// Delete
		tx.Delete(&readArtist)

		return nil
	})
	if err != nil {
		return
	}

}
