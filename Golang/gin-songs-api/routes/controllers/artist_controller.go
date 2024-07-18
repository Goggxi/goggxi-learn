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

// Create godoc
// @Summary Create a new artist
// @Description Create a new artist
// @Tags artists
// @Accept  json
// @Produce  json
// @Param artist body api.ArtistReq true "Artist request"
// @Success 201 {object} utils.SuccessResponse{data=api.Artist}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /artists [post]
func (ctrl *ArtisController) Create(c *gin.Context) {
	var req api.ArtistReq
	err := c.BindJSON(&req)
	if err != nil {
		utils.WriteErrorResponse(c, utils.BAD_REQUEST_MESSAGE, err.Error())
		return
	}

	err = ctrl.validate.Struct(req)
	if err != nil {
		utils.WriteErrorResponse(c, utils.VALIDATE_ERRORS_MESSAGE, err.Error())
		return
	}

	artist := &entity.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	data, err := ctrl.service.Create(c, artist)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.CREATED_MESSAGE, data)
}

// Update godoc
// @Summary Update an artist
// @Description Update an artist
// @Tags artists
// @Accept  json
// @Produce  json
// @Param id path string true "Artist ID"
// @Param artist body api.ArtistReq true "Artist request"
// @Success 200 {object} utils.SuccessResponse{data=api.Artist}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /artists/{id} [put]
func (ctrl *ArtisController) Update(c *gin.Context) {
	id := c.Param("id")

	var req api.ArtistReq
	err := c.BindJSON(&req)
	if err != nil {
		utils.WriteErrorResponse(c, utils.BAD_REQUEST_MESSAGE, err.Error())
		return
	}

	err = ctrl.validate.Struct(req)
	if err != nil {
		utils.WriteErrorResponse(c, utils.VALIDATE_ERRORS_MESSAGE, err.Error())
		return
	}

	artist := &entity.Artist{
		Name: req.Name,
		Bio:  req.Bio,
	}

	data, err := ctrl.service.Update(c, id, artist)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.UPDATED_MESSAGE, data)
}

// FindAll godoc
// @Summary Get all artists
// @Description Get all artists
// @Tags artists
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessResponse{data=[]api.Artist}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /artists [get]
func (ctrl *ArtisController) FindAll(c *gin.Context) {
	data, err := ctrl.service.FindAll(c)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// FindById godoc
// @Summary Get an artist by ID
// @Description Get an artist by ID
// @Tags artists
// @Accept  json
// @Produce  json
// @Param id path string true "Artist ID"
// @Success 200 {object} utils.SuccessResponse{data=api.Artist}
// @Failure 404 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /artists/{id} [get]
func (ctrl *ArtisController) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := ctrl.service.FindById(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// Delete godoc
// @Summary Delete an artist by ID
// @Description Delete an artist by ID
// @Tags artists
// @Accept  json
// @Produce  json
// @Param id path string true "Artist ID"
// @Success 200 {object} utils.SuccessResponse{data=nil}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /artists/{id} [delete]
func (ctrl *ArtisController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.Delete(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.DELETED_MESSAGE, nil)
}

func NewArtistController(service services.ArtistService) *ArtisController {
	return &ArtisController{service, validator.New()}
}
