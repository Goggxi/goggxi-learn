package controllers

import (
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
	"gin-songs-api/services"
	"gin-songs-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AlbumController struct {
	service  services.AlbumService
	validate *validator.Validate
}

// Create godoc
// @Summary Create a new album
// @Description Create a new album
// @Tags albums
// @Accept  json
// @Produce  json
// @Param album body api.AlbumReq true "Album request"
// @Success 201 {object} utils.SuccessResponse{data=api.Album}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /albums [post]
func (ctrl *AlbumController) Create(c *gin.Context) {
	var req api.AlbumReq
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

	album := &entity.Album{
		Title:       req.Title,
		Genre:       req.Genre,
		ArtistID:    req.ArtistID,
		ReleaseDate: req.ReleaseDate,
	}

	data, err := ctrl.service.Create(c, album)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.CREATED_MESSAGE, data)
}

// Update godoc
// @Summary Update an album
// @Description Update an album
// @Tags albums
// @Accept  json
// @Produce  json
// @Param id path string true "Album ID"
// @Param album body api.AlbumReq true "Album request"
// @Success 200 {object} utils.SuccessResponse{data=api.Album}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /albums/{id} [put]
func (ctrl *AlbumController) Update(c *gin.Context) {
	id := c.Param("id")

	var req api.AlbumReq
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

	album := &entity.Album{
		ID:          id,
		Title:       req.Title,
		Genre:       req.Genre,
		ArtistID:    req.ArtistID,
		ReleaseDate: req.ReleaseDate,
	}

	data, err := ctrl.service.Update(c, album)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.UPDATED_MESSAGE, data)
}

// FindAll godoc
// @Summary Get all albums
// @Description Get all albums
// @Tags albums
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessResponse{data=[]api.Album}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /albums [get]
func (ctrl *AlbumController) FindAll(c *gin.Context) {
	data, err := ctrl.service.FindAll(c)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// FindById godoc
// @Summary Get an album by ID
// @Description Get an album by ID
// @Tags albums
// @Accept  json
// @Produce  json
// @Param id path string true "Album ID"
// @Success 200 {object} utils.SuccessResponse{data=api.Album}
// @Failure 404 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /albums/{id} [get]
func (ctrl *AlbumController) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := ctrl.service.FindById(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// Delete godoc
// @Summary Delete an album by ID
// @Description Delete an album by ID
// @Tags albums
// @Accept  json
// @Produce  json
// @Param id path string true "Album ID"
// @Success 200 {object} utils.SuccessResponse{data=nil}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /albums/{id} [delete]
func (ctrl *AlbumController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.Delete(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.DELETED_MESSAGE, nil)
}

func NewAlbumController(service services.AlbumService) *AlbumController {
	return &AlbumController{service, validator.New()}
}
