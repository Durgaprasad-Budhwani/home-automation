package controllers

import (
	"net/http"

	"github.com/durgaprasad-budhwani/home-automation/backend/models"
	"github.com/durgaprasad-budhwani/home-automation/backend/services"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	service services.StatusService
}

func NewStatusController(service services.StatusService) StatusController {
	return StatusController{
		service,
	}
}

func (c StatusController) GetStatus(context *gin.Context) {
	schedulers, err := c.service.GetAll(10, 0)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, schedulers)
}

func (c StatusController) AddStatus(context *gin.Context) {
	var status models.Status
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

func (c StatusController) UpdateStatus(context *gin.Context) {
	var status models.Status
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
