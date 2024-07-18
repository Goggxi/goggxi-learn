package routes

import (
	"gin-songs-api/repositories"
	"gin-songs-api/routes/controllers"
	"gin-songs-api/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, pool *pgxpool.Pool) {
	// Repositories
	artistRepository := repositories.NewArtistRepository()

	// Services
	artistService := services.NewArtistService(artistRepository, pool)

	// Controllers
	artistController := controllers.NewArtistController(artistService)

	// Docs
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	// Path: /api/v1
	api := r.Group("/api/v1")

	// Path: /api/v1/artists
	artists := api.Group("/artists")
	{
		artists.POST("/", artistController.Create)
		artists.PUT("/:id", artistController.Update)
		artists.GET("/", artistController.FindAll)
		artists.GET("/:id", artistController.FindById)
		artists.DELETE("/:id", artistController.Delete)
	}

}
