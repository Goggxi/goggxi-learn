package controllers

import (
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
	"gin-songs-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArtisController struct {
	service  services.ArtistService
	validate *validator.Validate
}

func (ctrl *ArtisController) Create(c *gin.Context) {
	var req api.ArtistReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.validate.Struct(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	artist := &entity.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	data, err := ctrl.service.Create(c, artist)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, data)
}

func NewArtistController(service services.ArtistService) *ArtisController {
	return &ArtisController{service, validator.New()}
}
