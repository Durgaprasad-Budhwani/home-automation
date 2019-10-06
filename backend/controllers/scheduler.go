package controllers

import (
	"net/http"

	"github.com/durgaprasad-budhwani/home-automation/backend/models"
	"github.com/durgaprasad-budhwani/home-automation/backend/services"

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

func (c SchedulerController) GetSchedules(context *gin.Context) {
	schedulers, err := c.service.GetAll(10, 0)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, schedulers)
}

func (c SchedulerController) AddSchedule(context *gin.Context) {
	var status models.Schedule
	err := context.BindJSON(&status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Save(&status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, status)
}

func (c SchedulerController) UpdateSchedule(context *gin.Context) {
	var status models.Schedule
	err := context.BindJSON(&status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Update(status.ID, &status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, status)
}
