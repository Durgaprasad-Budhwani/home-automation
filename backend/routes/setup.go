package routes

import (
	"github.com/durgaprasad-budhwani/home-automation/backend/controllers"
	"github.com/durgaprasad-budhwani/home-automation/backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	//// use compression
	r.Use(gzip.Gzip(gzip.BestSpeed))

	// enable CORS for all domain
	r.Use(cors.Default())

	// services
	schedulerService := services.NewSchedulerService(db)
	schedulerController := controllers.NewSchedulerController(schedulerService)

	// Ping test
	r.GET("/schedulers/", schedulerController.GetSchedulers)
	r.POST("/schedulers", schedulerController.AddScheduler)
	r.PUT("/schedulers/:id", schedulerController.UpdateScheduler)

	return r
}
