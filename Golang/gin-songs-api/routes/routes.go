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
	albumRepository := repositories.NewAlbumRepository()

	// Services
	artistService := services.NewArtistService(artistRepository, pool)
	albumService := services.NewAlbumService(albumRepository, artistRepository, pool)

	// Controllers
	artistController := controllers.NewArtistController(artistService)
	albumController := controllers.NewAlbumController(albumService)

	// Docs
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	// Path: /api/v1
	api := r.Group("/api/v1")

	// Path: /api/v1/artists
	artists := api.Group("/artists")
	{
		artists.POST("", artistController.Create)
		artists.PUT("/:id", artistController.Update)
		artists.GET("", artistController.FindAll)
		artists.GET("/:id", artistController.FindById)
		artists.DELETE("/:id", artistController.Delete)
	}

	// Path: /api/v1/albums
	albums := api.Group("/albums")
	{
		albums.POST("", albumController.Create)
		albums.PUT("/:id", albumController.Update)
		albums.GET("", albumController.FindAll)
		albums.GET("/:id", albumController.FindById)
		albums.DELETE("/:id", albumController.Delete)
	}

}
