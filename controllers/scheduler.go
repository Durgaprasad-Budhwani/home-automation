package controllers

import (
	"net/http"

	"github.com/durgaprasad-budhwani/home-automation/models"
	"github.com/durgaprasad-budhwani/home-automation/services"

	"github.com/gin-gonic/gin"
)

type SchedulerController struct {
	service services.SchedulerService
}

func NewSchedulerController(service services.SchedulerService) SchedulerController {
	return SchedulerController{
		service,
	}
}

func (c SchedulerController) GetSchedulers(context *gin.Context) {
	schedulers, err := c.service.GetAll(10, 0)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, schedulers)
}

func (c SchedulerController) AddScheduler(context *gin.Context) {
	var scheduler models.Scheduler
	err := context.BindJSON(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Save(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, scheduler)
}

func (c SchedulerController) UpdateScheduler(context *gin.Context) {
	var scheduler models.Scheduler
	err := context.BindJSON(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Update(scheduler.ID, &scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, scheduler)
}
