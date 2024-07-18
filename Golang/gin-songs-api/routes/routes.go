package routes

import (
	"gin-songs-api/repositories"
	"gin-songs-api/routes/controllers"
	"gin-songs-api/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

func SetupRoutes(r *gin.Engine, pool *pgxpool.Pool) {
	// Repositories
	artistRepository := repositories.NewArtistRepository()

	// Services
	artistService := services.NewArtistService(artistRepository, pool)

	// Controllers
	artistController := controllers.NewArtistController(artistService)

	// Welcome route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	api.POST("/artists", artistController.Create)
}
