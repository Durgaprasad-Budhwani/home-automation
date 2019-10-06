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
	schedulerService := services.NewSlotService(db)
	schedulerController := controllers.NewSlotController(schedulerService)

	statusService := services.NewSchedulerService(db)
	statusController := controllers.NewSchedulerController(statusService)

	// Ping test
	r.GET("/schedulers/", schedulerController.GetSlots)
	r.POST("/schedulers", schedulerController.AddSlot)
	r.PUT("/schedulers/:id", schedulerController.UpdateSlot)

	r.GET("/status/", statusController.GetSchedulers)
	r.POST("/status", statusController.AddSchedule)
	r.PUT("/status/:id", statusController.UpdateSchedule)

	return r
}
