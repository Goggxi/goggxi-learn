package controllers

import (
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
	"gin-songs-api/services"
	"gin-songs-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SongController struct {
	service  services.SongService
	validate *validator.Validate
}

// Create godoc
// @Summary Create a new song
// @Description Create a new song
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body api.SongReq true "Song request"
// @Success 201 {object} utils.SuccessResponse{data=api.Song}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /songs [post]
func (ctrl *SongController) Create(c *gin.Context) {
	var req api.SongReq
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
		ID: req.AlbumID,
	}

	song := &entity.Song{
		Title:       req.Title,
		Album:       *album,
		Duration:    req.Duration,
		ReleaseDate: req.ReleaseDate,
	}

	data, err := ctrl.service.Create(c, song)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.CREATED_MESSAGE, data)
}

// Update godoc
// @Summary Update a song
// @Description Update a song
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path string true "Song ID"
// @Param song body api.SongReq true "Song request"
// @Success 200 {object} utils.SuccessResponse{data=api.Song}
// @Failure 400 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /songs/{id} [put]
func (ctrl *SongController) Update(c *gin.Context) {
	id := c.Param("id")

	var req api.SongReq
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
		ID: req.AlbumID,
	}

	song := &entity.Song{
		ID:          id,
		Title:       req.Title,
		Album:       *album,
		Duration:    req.Duration,
		ReleaseDate: req.ReleaseDate,
	}

	data, err := ctrl.service.Update(c, song)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.UPDATED_MESSAGE, data)
}

// FindAll godoc
// @Summary Get all songs
// @Description Get all songs
// @Tags songs
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessResponse{data=[]api.Song}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /songs [get]
func (ctrl *SongController) FindAll(c *gin.Context) {
	data, err := ctrl.service.FindAll(c)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// FindById godoc
// @Summary Get a song by ID
// @Description Get a song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path string true "Song ID"
// @Success 200 {object} utils.SuccessResponse{data=api.Song}
// @Failure 404 {object} utils.ErrorResponse{error=string}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /songs/{id} [get]
func (ctrl *SongController) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := ctrl.service.FindById(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.SUCCESS_MESSAGE, data)
}

// Delete godoc
// @Summary Delete a song by ID
// @Description Delete a song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path string true "Song ID"
// @Success 200 {object} utils.SuccessResponse{data=nil}
// @Failure 500 {object} utils.ErrorResponse{error=string}
// @Router /songs/{id} [delete]
func (ctrl *SongController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.Delete(c, id)
	if err != nil {
		utils.WriteErrorResponse(c, utils.INTERNAL_SERVER_ERROR_MESSAGE, err.Error())
		return
	}

	utils.WriteSuccessResponse(c, utils.DELETED_MESSAGE, nil)
}

func NewSongController(service services.SongService) *SongController {
	return &SongController{service, validator.New()}
}
