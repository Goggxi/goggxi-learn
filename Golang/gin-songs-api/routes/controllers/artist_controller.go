package controllers

import (
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
	"gin-songs-api/services"
	"gin-songs-api/utils"
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
		utils.ErrorResponse(c, utils.BAD_REQUEST_MESSAGE, err.Error())
		return
	}

	err = ctrl.validate.Struct(req)
	if err != nil {
		utils.ErrorResponse(c, utils.VALIDATE_ERRORS_MESSAGE, err.Error())
		return
	}

	artist := &entity.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	data, err := ctrl.service.Create(c, artist)
	if err != nil {
		utils.ErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.SuccessResponse(c, utils.CREATED_MESSAGE, data)
}

func (ctrl *ArtisController) Update(c *gin.Context) {
	id := c.Param("id")

	var req api.ArtistReq
	err := c.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(c, utils.BAD_REQUEST_MESSAGE, err.Error())
		return
	}

	err = ctrl.validate.Struct(req)
	if err != nil {
		utils.ErrorResponse(c, utils.VALIDATE_ERRORS_MESSAGE, err.Error())
		return
	}

	artist := &entity.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	data, err := ctrl.service.Update(c, id, artist)
	if err != nil {
		utils.ErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.SuccessResponse(c, utils.UPDATED_MESSAGE, data)
}

func (ctrl *ArtisController) FindAll(c *gin.Context) {
	data, err := ctrl.service.FindAll(c)
	if err != nil {
		utils.ErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.SuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

func (ctrl *ArtisController) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := ctrl.service.FindById(c, id)
	if err != nil {
		utils.ErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.SuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

func (ctrl *ArtisController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.Delete(c, id)
	if err != nil {
		utils.ErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.SuccessResponse(c, utils.DELETED_MESSAGE, nil)
}

func NewArtistController(service services.ArtistService) *ArtisController {
	return &ArtisController{service, validator.New()}
}
